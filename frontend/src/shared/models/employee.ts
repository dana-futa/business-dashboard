export interface Employee {
    employeeId: number;
    firstName: string;
    lastName: string;
    email: string;
    title: string;
    managerId: number;
    reports: Array<Employee>;
    isActive: boolean;
}
