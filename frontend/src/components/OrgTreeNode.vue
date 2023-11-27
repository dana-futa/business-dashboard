<script lang="ts">
import type { Employee } from '@/shared/models/employee';
import { defineComponent } from 'vue'

export default defineComponent({
    name: 'OrgTreeNode',
    props: {
        employee: {
            type: Object as () => Employee,
            required: true
        },
  },
})
</script>

<template>
    <div class="org-tree-container">
        <div class="org-tree-node">
            <div>First Name: {{ employee.firstName }}</div>
            <div>Last Name: {{ employee.lastName }}</div>
            <div>Employee ID: {{ employee.employeeId }}</div>
            <div>Email: {{ employee.email }}</div>
            <div>Title: {{ employee.title }}</div>
            <div>Is Active: {{ employee.isActive }}</div>
            <div>Manager ID: {{ employee.managerId }}</div>
        </div>
        <div v-if="employee.reports && employee.reports.length > 0" class="children">
            <OrgTreeNode v-for="report in employee.reports" :key="report.employeeId" :employee="report" />
        </div>
    </div>
</template>

<style>
.org-tree-container {
    flex: 0;
    margin: 1rem;
}

.org-tree-node {
    border: 1px solid black;
    border-radius: 5px;
    width: 20rem;
    margin: auto;
    padding: 1rem;
}

.children {
    display: flex;
    flex-direction: row;
}
</style>
