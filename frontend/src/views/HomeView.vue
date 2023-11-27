<script setup lang="ts">
import type { Employee } from '@/shared/models/employee';
import OrgTree from '../components/OrgTree.vue';

const employeeTree: Employee =
    {
        employeeId: 0, 
        firstName: "Bob", 
        lastName: "Stevens", 
        email: "bob.stevens@business.com", 
        title: "CEO", 
        managerId: -1, 
        isActive: true,
        reports: [
            {
                employeeId: 1, 
                firstName: "Mary", 
                lastName: "Jane", 
                email: "mary.jane@business.com", 
                title: "Lead Engineer", 
                managerId: 0, 
                isActive: true,
                reports: [
                    {
                        employeeId: 3, 
                        firstName: "David --> reports to mary", 
                        lastName: "Gomez", 
                        email: "david.gomez@business.com", 
                        title: "Engineer", 
                        managerId: 1, 
                        isActive: true,
                        reports: [
                            {
                                employeeId: 30, 
                                firstName: "Jerry --> reports to David", 
                                lastName: "Sanchez", 
                                email: "jerry.sanchez@business.com", 
                                title: "Engineer", 
                                managerId: 3, 
                                isActive: true,
                                reports: [],
                            },
                            {
                                employeeId: 31, 
                                firstName: "Gabby --> reports to David", 
                                lastName: "Lemon", 
                                email: "gabby.lemon@business.com", 
                                title: "Engineer", 
                                managerId: 3, 
                                isActive: true,
                                reports: [],
                            },
                        ],
                    },
                    {
                        employeeId: 21, 
                        firstName: "Sally", 
                        lastName: "Michaelson", 
                        email: "sally.michaelson@business.com", 
                        title: "Engineer 1", 
                        managerId: 0, 
                        isActive: true,
                        reports: [],
                    },
                ]
            },
            {
                employeeId: 2, 
                firstName: "John", 
                lastName: "Hope", 
                email: "john.hope@business.com", 
                title: "Engineer 1", 
                managerId: 0, 
                isActive: true,
                reports: [],
            },
            {
                employeeId: 42, 
                firstName: "Jessica", 
                lastName: "Davis", 
                email: "jessica.davis@business.com", 
                title: "Engineer 1", 
                managerId: 0, 
                isActive: true,
                reports: [],
            },
        ],
};

console.log('employeeTree', employeeTree);
</script>

<template>
    <main>
        <OrgTree :root="employeeTree" />
    </main>
</template>
