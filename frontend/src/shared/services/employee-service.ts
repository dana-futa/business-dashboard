import axios from 'axios'; // , { type AxiosResponse }
// import type { Employee } from '../models/employee';

const API_URL = 'http://localhost:8080/api/v1';

export function getTestEmployees() {
    return axios.get(`${API_URL}/test-employees`)
        .then((response) => {
            console.log(response)
        });
}
