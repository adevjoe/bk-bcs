<template>
  <section class="create-version">
    <bk-button
      v-if="versionData.status.publish_status === 'partial_released'"
      v-cursor="{ active: !props.hasPerm }"
      theme="primary"
      :class="['trigger-button', { 'bk-button-with-no-perm': !props.hasPerm }]"
      :disabled="props.permCheckLoading"
      @click="handleBtnClick">
      {{ t('调整分组上线') }}
    </bk-button>
    <Teleport to="body">
      <VersionLayout v-if="isSelectGroupPanelOpen">
        <template #header>
          <section class="header-wrapper">
            <span class="header-name" @click="handlePanelClose">
              <ArrowsLeft class="arrow-left" />
              <span class="service-name">{{ appData.spec.name }}</span>
            </span>
            <AngleRight class="arrow-right" />
            {{ t('调整分组上线') }}：{{ versionData.spec.name }}
          </section>
        </template>
        <select-group
          ref="selectGroupRef"
          :loading="versionListLoading || groupListLoading"
          :group-list="treeNodeGroups"
          :version-list="versionList"
          :groups="groups"
          :release-type="releaseType"
          :released-groups="releasedGroups"
          :disable-select="disableSelect"
          @open-preview-version-diff="openPreviewVersionDiff"
          @release-type-change="releaseType = $event"
          @change="groups = $event">
        </select-group>
        <template #footer>
          <section class="actions-wrapper">
            <bk-button
              v-bk-tooltips="{ content: t('请选择需要调整的分组'), disabled: isSelectedGroupsChanged }"
              class="publish-btn"
              theme="primary"
              :disabled="!isSelectedGroupsChanged"
              @click="handlePublishOrOpenDiff">
              {{ diffableVersionList.length ? t('对比并上线') : t('上线版本') }}
            </bk-button>
            <bk-button @click="handlePanelClose">取消</bk-button>
          </section>
        </template>
      </VersionLayout>
    </Teleport>
    <ConfirmDialog
      v-model:show="isConfirmDialogShow"
      :bk-biz-id="props.bkBizId"
      :app-id="props.appId"
      :group-list="treeNodeGroups"
      :version-list="versionList"
      :release-type="releaseType"
      :groups="groups"
      :released-groups="releasedGroups"
      @confirm="handleConfirm" />
    <PublishVersionDiff
      :bk-biz-id="props.bkBizId"
      :app-id="props.appId"
      :show="isDiffSliderShow"
      :current-version="versionData"
      :base-version-id="baseVersionId"
      :version-list="diffableVersionList"
      :current-version-groups="groupsPendingtoPublish"
      @publish="handleOpenPublishDialog"
      @close="isDiffSliderShow = false" />
  </section>
</template>
<script setup lang="ts">
  import { ref, computed, watch } from 'vue';
  import { useI18n } from 'vue-i18n';
  import { useRouter } from 'vue-router';
  import { ArrowsLeft, AngleRight } from 'bkui-vue/lib/icon';
  import { InfoBox } from 'bkui-vue';
  import BkMessage from 'bkui-vue/lib/message';
  import { storeToRefs } from 'pinia';
  import { getConfigVersionList } from '../../../../../api/config';
  import { getServiceGroupList } from '../../../../../api/group';
  import { IConfigVersion, IReleasedGroup } from '../../../../../../types/config';
  import useGlobalStore from '../../../../../store/global';
  import useServiceStore from '../../../../../store/service';
  import useConfigStore from '../../../../../store/config';
  import { IGroupToPublish, IGroupItemInService } from '../../../../../../types/group';
  import VersionLayout from '../config/components/version-layout.vue';
  import ConfirmDialog from './publish-version/confirm-dialog.vue';
  import SelectGroup from './publish-version/select-group/index.vue';
  import PublishVersionDiff from './publish-version-diff.vue';

  const { permissionQuery, showApplyPermDialog } = storeToRefs(useGlobalStore());
  const serviceStore = useServiceStore();
  const versionStore = useConfigStore();
  const { appData } = storeToRefs(serviceStore);
  const { versionData, publishedVersionId } = storeToRefs(versionStore);
  const { t } = useI18n();

  const props = defineProps<{
    bkBizId: string;
    appId: number;
    permCheckLoading: boolean;
    hasPerm: boolean;
  }>();

  const emit = defineEmits(['confirm']);

  const router = useRouter();
  const versionList = ref<IConfigVersion[]>([]);
  const versionListLoading = ref(true);
  const groupList = ref<IGroupItemInService[]>([]);
  const groupListLoading = ref(true);
  const treeNodeGroups = ref<IGroupToPublish[]>([]);
  const isSelectGroupPanelOpen = ref(false);
  const isDiffSliderShow = ref(false);
  const isConfirmDialogShow = ref(false);
  const releaseType = ref('select');
  const disableSelect = ref(false);
  const groups = ref<IGroupToPublish[]>([]);
  const baseVersionId = ref(0);
  const selectGroupRef = ref();

  // 当前版本已上线分组id集合
  const releasedGroups = computed(() => versionData.value.status.released_groups.map((group) => group.id));

  // 包含分组变更的版本，用来对比线上版本
  const diffableVersionList = computed(() => {
    const list = [] as IConfigVersion[];
    versionList.value.forEach((version) => {
      if (version.id === versionData.value.id) return; // 忽略当前上线版本
      version.status.released_groups.some((item) => {
        if (releasedGroups.value.includes(item.id)) return false; // 忽略已在当前版本上线的分组
        if (item.id === 0) {
          // 其他版本包含默认分组，且当前选中分组未上线
          return groups.value.some((g) => {
            if (g.release_id === 0 || g.id === 0) {
              list.push(version);
              return true;
            }
            return false;
          });
        }

        // 其他版本包含的分组在当前已选中的分组中
        if (groups.value.findIndex((g) => g.id === item.id) > -1) {
          list.push(version);
          return true;
        }
        return false;
      });
    });
    return list;
  });

  // 是否有调整分组
  const isSelectedGroupsChanged = computed(() => {
    return groups.value.length > 0 && groups.value.some((item) => !releasedGroups.value.includes(item.id));
  });

  const permissionQueryResource = computed(() => [
    {
      biz_id: props.bkBizId,
      basic: {
        type: 'app',
        action: 'publish',
        resource_id: props.appId,
      },
    },
  ]);

  // 待上线分组实例
  const groupsPendingtoPublish = computed(() => {
    const list: IReleasedGroup[] = [];
    groups.value.forEach((item) => {
      if (releasedGroups.value.includes(item.id)) return;
      const group = groupList.value.find((g) => g.group_id === item.id);
      if (group) {
        const { group_id, group_name, new_selector, old_selector, edited } = group;
        list.push({
          id: group_id,
          name: group_name,
          new_selector,
          old_selector,
          edited,
          uid: '',
          mode: '',
        });
      }
    });
    return list;
  });

  watch(
    () => versionData.value,
    () => {
      setReleaseData();
    },
  );

  // 判断是否需要对比上线版本
  const handlePublishOrOpenDiff = () => {
    if (diffableVersionList.value.length) {
      baseVersionId.value = diffableVersionList.value[0].id;
      isDiffSliderShow.value = true;
      return;
    }
    handleOpenPublishDialog();
  };

  // 获取所有已上线版本（已上线或灰度中）
  const getVersionList = async () => {
    try {
      versionListLoading.value = true;
      const res = await getConfigVersionList(props.bkBizId, props.appId, { start: 0, all: true });
      versionList.value = res.data.details.filter((item: IConfigVersion) => {
        return item.status.publish_status !== 'not_released';
      });
      versionListLoading.value = false;
    } catch (e) {
      console.error(e);
    }
  };

  // 获取所有分组，并组装tree组件节点需要的数据
  const getAllGroupData = async () => {
    groupListLoading.value = true;
    const res = await getServiceGroupList(props.bkBizId, appData.value.id as number);
    groupList.value = res.details;
    treeNodeGroups.value = res.details.map((group: IGroupItemInService) => {
      const { group_id, group_name, release_id, release_name } = group;
      const selector = group.new_selector;
      const rules = selector.labels_and || selector.labels_or || [];
      return { id: group_id, name: group_name, release_id, release_name, rules };
    });

    groupListLoading.value = false;
  };

  const handleBtnClick = () => {
    if (props.hasPerm) {
      getVersionList();
      getAllGroupData();
      setReleaseData();
      openSelectGroupPanel();
    } else {
      permissionQuery.value = { resources: permissionQueryResource.value };
      showApplyPermDialog.value = true;
    }
  };

  // 打开选择分组面板
  const openSelectGroupPanel = () => {
    isSelectGroupPanelOpen.value = true;
    groups.value = versionData.value.status.released_groups.map((group) => {
      const { id, name } = group;
      const selector = group.new_selector;
      const rules = selector?.labels_and || [];
      return {
        id,
        name,
        release_id: versionData.value.id,
        release_name: versionData.value.spec.name,
        disabled: true,
        rules,
      };
    });
  };

  const setReleaseData = () => {
    if (versionData.value.status.released_groups.some((group) => group.id === 0)) {
      releaseType.value = 'all';
      disableSelect.value = true;
    } else {
      releaseType.value = 'select';
      disableSelect.value = false;
    }
  };

  // 打开上线版本确认弹窗
  const handleOpenPublishDialog = () => {
    if (groups.value.length === 0) {
      BkMessage({ theme: 'error', message: '请选择分组实例' });
      return;
    }
    isConfirmDialogShow.value = true;
  };

  // 选择分组面板上线预览版本对比
  const openPreviewVersionDiff = (id: number) => {
    baseVersionId.value = id;
    isDiffSliderShow.value = true;
  };

  // 上线确认
  const handleConfirm = (havePull: boolean) => {
    isDiffSliderShow.value = false;
    publishedVersionId.value = versionData.value.id;
    handlePanelClose();
    emit('confirm');
    if (havePull) {
      InfoBox({
        infoType: 'success',
        'ext-cls': 'info-box-style',
        title: t('调整分组上线成功'),
        dialogType: 'confirm',
      });
    } else {
      InfoBox({
        infoType: 'success',
        'ext-cls': 'info-box-style',
        title: t('调整分组上线成功'),
        confirmText: t('配置客户端'),
        cancelText: t('稍后再说'),
        onConfirm: () => {
          const routeData = router.resolve({
            name: 'configuration-example',
            params: { spaceId: props.bkBizId, appId: props.appId },
          });
          window.open(routeData.href, '_blank');
        },
      });
    }
  };

  // 关闭选择分组面板
  const handlePanelClose = () => {
    releaseType.value = 'all';
    isSelectGroupPanelOpen.value = false;
    groups.value = [];
  };
</script>
<style lang="scss" scoped>
  .trigger-button {
    margin-left: 8px;
  }
  .header-wrapper {
    display: flex;
    align-items: center;
    padding: 0 24px;
    height: 100%;
    font-size: 12px;
    line-height: 1;
  }
  .header-name {
    display: flex;
    align-items: center;
    font-size: 12px;
    color: #3a84ff;
    cursor: pointer;
  }
  .arrow-left {
    font-size: 26px;
    color: #3884ff;
  }
  .arrow-right {
    font-size: 24px;
    color: #c4c6cc;
  }
  .actions-wrapper {
    display: flex;
    align-items: center;
    padding: 0 24px;
    height: 100%;
    .publish-btn {
      margin-right: 8px;
    }
    .bk-button {
      min-width: 88px;
    }
  }
  .version-selector {
    display: flex;
    align-items: center;
    height: 100%;
    padding: 0 24px;
    font-size: 12px;
  }
</style>
