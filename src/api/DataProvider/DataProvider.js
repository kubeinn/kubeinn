import axios from 'axios';
import { CreateProjectByPilgrim, DeleteProjectByPilgrim, DeleteProjectsByPilgrim } from '../Pilgrim/Pilgrim';

// Production
// const dataProviderUrl = window._env_.KUBEINN_POSTGREST_URL;
// Local
const dataProviderUrl = process.env.REACT_APP_KUBEINN_POSTGREST_URL;

export default {
    getList: (resource, params) => {
        const { page, perPage } = params.pagination;
        const { field, order } = params.sort;
        const url = `${dataProviderUrl}/${resource}`;

        return axios({
            url: url,
            method: 'GET',
            headers: {
                'Prefer': 'count=exact',
                'Authorization': getCookie("Authorization"),
                'Content-Type': 'application/json'
            },
            params: {
                offset: JSON.stringify((page - 1) * perPage),
                limit: JSON.stringify(perPage),
                order: field + "." + order.toLowerCase(),
            },
            timeout: 5000,
            responseType: 'json',
            responseEncoding: 'utf8',
        }).then(response => ({
            data: response.data,
            total: parseInt(response.headers['content-range'].split('/').pop(), 10),
        }));
    },

    create: (resource, params) => {
        const url = `${dataProviderUrl}/${resource}`;
        const { createProject, ...record } = params.data;

        let updateDatabase = axios({
            url: url,
            method: 'POST',
            headers: {
                'Authorization': getCookie("Authorization"),
            },
            data: record,
            timeout: 5000,
            responseType: 'json',
            responseEncoding: 'utf8',
        }).then(({ json }) => ({
            data: { ...params.data },
        }));

        if (createProject) {
            console.log("createProject is true.")
            return CreateProjectByPilgrim(params).then(
                updateDatabase
            );
        }
        return updateDatabase;
    },

    delete: (resource, params) => {
        const url = `${dataProviderUrl}/${resource}`;
        const paramsId = "eq." + params.id;

        let updateDatabase = axios({
            url: url,
            method: 'DELETE',
            headers: {
                'Authorization': getCookie("Authorization"),
            },
            params: {
                id: paramsId,
            },
            timeout: 5000,
            responseType: 'json',
            responseEncoding: 'utf8',
        }).then(response => ({
            data: response.data,
        }));

        console.log(resource)
        if (resource == 'projects') {
            console.log("deleteProject is true.")
            return DeleteProjectByPilgrim(params).then(
                updateDatabase
            );
        }
        return updateDatabase;
    },

    deleteMany: (resource, params) => {
        const url = `${dataProviderUrl}/${resource}`;

        var arrayLength = params.ids.length;
        var paramsId = "(id.eq." + params.ids[0];
        for (var i = 1; i < arrayLength; i++) {
            console.log(paramsId[i]);
            paramsId = paramsId + ",id.eq." + params.ids[i];
        }
        paramsId = paramsId + ")";

        let updateDatabase = axios({
            url: url,
            method: 'DELETE',
            headers: {
                'Authorization': getCookie("Authorization"),
            },
            params: {
                or: paramsId,
            },
            timeout: 5000,
            responseType: 'json',
            responseEncoding: 'utf8',
        }).then(response => ({
            data: response.data,
        }));

        console.log(resource)
        if (resource == 'projects') {
            console.log("deleteProjects is true.")
            return DeleteProjectsByPilgrim(params).then(
                updateDatabase
            );
        }
        return updateDatabase;
    },
};

function setCookie(name, value, days) {
    var expires = "";
    if (days) {
        var date = new Date();
        date.setTime(date.getTime() + (days * 24 * 60 * 60 * 1000));
        expires = "; expires=" + date.toUTCString();
    }
    document.cookie = name + "=" + (value || "") + expires + "; path=/";
}
function getCookie(name) {
    const value = `; ${document.cookie}`;
    const parts = value.split(`; ${name}=`);
    if (parts.length === 2) return parts.pop().split(';').shift();
}
function eraseCookie(name) {
    document.cookie = name + '=; Path=/; Expires=Thu, 01 Jan 1970 00:00:01 GMT;';
}