import axios from 'axios';
import { CreateProjectByPilgrim, DeleteProjectByPilgrim, DeleteProjectsByPilgrim } from '../Pilgrim/Pilgrim';

var dataProviderUrl;
if (!process.env.NODE_ENV || process.env.NODE_ENV === 'development') {
    // dev code
    dataProviderUrl = process.env.REACT_APP_KUBEINN_POSTGREST_URL;
} else {
    // production code
    dataProviderUrl = window._env_.KUBEINN_POSTGREST_URL;
}
console.log(dataProviderUrl)

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

    create: async (resource, params) => {
        const url = `${dataProviderUrl}/${resource}`;
        const { createProject, ...record } = params.data;

        async function updateDatabase() {
            var data = await axios({
                url: url,
                method: 'POST',
                headers: {
                    'Authorization': getCookie("Authorization"),
                    'Prefer': 'return=representation',
                },
                data: record,
                timeout: 5000,
                responseType: 'json',
                responseEncoding: 'utf8',
            }).then((response) => ({
                data: { ...params.data, id: response.id },
            }));
            return data;
        }

        if (createProject) {
            console.log("createProject is true.");
            var response = await CreateProjectByPilgrim(params);
            console.log(response)
            record["kube_configuration"] = response.data.kubecfg;
            return updateDatabase();
        }

        return updateDatabase();
    },

    delete: (resource, params) => {
        const url = `${dataProviderUrl}/${resource}`;
        const paramsId = "eq." + params.id;

        async function updateDatabase() {
            var data = await axios({
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
            return data;
        };

        console.log(resource);
        if (resource == 'projects') {
            console.log("deleteProject is true.");
            DeleteProjectByPilgrim(params);
            return updateDatabase();
        }
        return updateDatabase();
    },

    deleteMany: async (resource, params) => {
        const url = `${dataProviderUrl}/${resource}`;

        var arrayLength = params.ids.length;
        var paramsId = "(id.eq." + params.ids[0];
        for (var i = 1; i < arrayLength; i++) {
            console.log(paramsId[i]);
            paramsId = paramsId + ",id.eq." + params.ids[i];
        }
        paramsId = paramsId + ")";

        async function updateDatabase() {
            var data = await axios({
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
            return data;
        };

        console.log(resource)
        if (resource == 'projects') {
            console.log("deleteProjects is true.");
            await DeleteProjectsByPilgrim(params);
            return updateDatabase();
        }
        return updateDatabase();
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