import request from '../utils/request';

export const fetchData = (query) => {
    return request({
        url: 'https://www.easy-mock.com/mock/592501a391470c0ac1fab128/ms/table/list',
        method: 'post',
        data: query
    })
};
export const userData = (query) => {
    return request({
        url: 'v1/getUsers',
        method: 'get',
        params: query
    })
};
export const updateUser = (query) => {
    return request({
        url: 'v1/updateUser',
        method: 'post',
        data: query
    })
};
export const addUser = (query) => {
    return request({
        url: 'v1/insertUser',
        method: 'post',
        data: query
    })
};
export const deleteUser = (query) => {
    return request({
        url: 'v1/deleteUser',
        method: 'post',
        data: query
    })
};