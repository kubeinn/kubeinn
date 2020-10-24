import { fetchUtils } from 'react-admin';
import { stringify } from 'query-string';
import axios from 'axios';

// Production
// const dataProviderUrl = window._env_.KUBEINN_POSTGREST_URL;
// Local
const dataProviderUrl = process.env.REACT_APP_KUBEINN_POSTGREST_URL;

// Instantiate httpClient
const httpClient = fetchUtils.fetchJson;

export default {
    getList: (resource, params) => {
        const { page, perPage } = params.pagination;
        const { field, order } = params.sort;
        const url = `${dataProviderUrl}/${resource}`;

        return axios.get(url, {
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
        })
            .then(response => ({
                data: response.data,
                total: parseInt(response.headers['content-range'].split('/').pop(), 10),
            }));
    },

    create: (resource, params) =>
        httpClient(`${dataProviderUrl}/${resource}`, {
            method: 'POST',
            body: JSON.stringify(params.data),
        }).then(({ json }) => ({
            data: { ...params.data },
        })),

    delete: (resource, params) => {
        const url = `${dataProviderUrl}/${resource}`;
        const paramsId = "eq." + params.id;

        return axios.delete(url, {
            headers: {},
            params: {
                id: paramsId,
            },
            timeout: 1000,
            responseType: 'json',
            responseEncoding: 'utf8',
        })
            .then(response => ({
                data: response.data,
            }));
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

        return axios.delete(url, {
            headers: {},
            params: {
                or: paramsId,
            },
            timeout: 1000,
            responseType: 'json',
            responseEncoding: 'utf8',
        })
            .then(response => ({
                data: response.data,
            }));
    },

    // {
    //     const query = {
    //         filter: JSON.stringify({ id: params.ids }),
    //     };
    //     return httpClient(`${dataProviderUrl}/${resource}?${stringify(query)}`, {
    //         method: 'DELETE',
    //         body: JSON.stringify(params.data),
    //     }).then(({ json }) => ({ data: json }));
    // },

    // getOne: (resource, params) =>
    //     httpClient(`${dataProviderUrl}/${resource}/${params.id}`)
    //     .then(({ json }) => ({
    //         data: json,
    //     })),

    // getMany: (resource, params) => {
    //     const query = {
    //         filter: JSON.stringify({ id: params.ids }),
    //     };
    //     const url = `${dataProviderUrl}/${resource}?${stringify(query)}`;
    //     return httpClient(url).then(({ json }) => ({ data: json }));
    // },

    // getManyReference: (resource, params) => {
    //     const { page, perPage } = params.pagination;
    //     const { field, order } = params.sort;
    //     const query = {
    //         sort: JSON.stringify([field, order]),
    //         range: JSON.stringify([(page - 1) * perPage, page * perPage - 1]),
    //         filter: JSON.stringify({
    //             ...params.filter,
    //             [params.target]: params.id,
    //         }),
    //     };
    //     const url = `${dataProviderUrl}/${resource}?${stringify(query)}`;

    //     return httpClient(url).then(({ headers, json }) => ({
    //         data: json,
    //         total: parseInt(headers.get('content-range').split('/').pop(), 10),
    //     }));
    // },

    // update: (resource, params) =>
    //     httpClient(`${dataProviderUrl}/${resource}/${params.id}`, {
    //         method: 'PUT',
    //         body: JSON.stringify(params.data),
    //     }).then(({ json }) => ({ data: json })),

    // updateMany: (resource, params) => {
    //     const query = {
    //         filter: JSON.stringify({ id: params.ids }),
    //     };
    //     return httpClient(`${dataProviderUrl}/${resource}?${stringify(query)}`, {
    //         method: 'PUT',
    //         body: JSON.stringify(params.data),
    //     }).then(({ json }) => ({ data: json }));
    // },
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