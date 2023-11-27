<script lang="ts">
import type { Employee } from '@/shared/models/employee';
import { defineComponent } from 'vue'
import OrgTreeNode from './OrgTreeNode.vue';

export default defineComponent({
    name: 'OrgTree',
    props: {
        root: {
            type: Object as () => Employee,
            required: true
        },
    },
    components: { OrgTreeNode }
})
</script>

<template>
    <div class="org-tree">
        <OrgTreeNode :employee="root" />
    </div>
</template>

<style>
.org-tree {
    display: flex;
}
</style>
