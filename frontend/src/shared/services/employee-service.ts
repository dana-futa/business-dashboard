import axios from 'axios'; // , { type AxiosResponse }
// import type { Employee } from '../models/employee';

const API_URL = 'http://localhost:8080/api/v1';
const axiosClient = axios.create({
	baseURL: API_URL,
});



export async function getTestEmployees(): {err: any, response: Array<Employee>} {
    return await axiosClient.get(`/test-employees`)
        .then((response) => {
            console.log('1', response);
            return {err: null, response: response.data.data}
        })
        .catch((err) => {
            console.error(err);
            return {err, response: []}
        });
}

export async function getTestEmployees2(): {err: any, response: Array<Employee>} {
    return await axiosClient.get<any>(`/all-active-employees`)
        .then((response) => {
            console.log('2', response)
            return {err: null, response: response.data.data}
        })
        .catch((err) => {
            console.error(err);
            return {err, response: []}
        });
}
