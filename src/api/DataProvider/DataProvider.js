import { fetchUtils } from 'react-admin';
import { stringify } from 'query-string';
import axios from 'axios';

// Production
// const apiUrl = window._env_.KUBEINN_POSTGREST_URL;

// Local
const apiUrl = process.env.REACT_APP_KUBEINN_POSTGREST_URL;

const httpClient = fetchUtils.fetchJson;

export default {
    getList: (resource, params) => {
        const { page, perPage } = params.pagination;
        const { field, order } = params.sort;
        const url = `${apiUrl}/${resource}`;

        return axios.get(url, {
            headers: { 'Prefer': 'count=exact' },
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
        httpClient(`${apiUrl}/${resource}`, {
            method: 'POST',
            body: JSON.stringify(params.data),
        }).then(({ json }) => ({
            data: { ...params.data },
        })),

    delete: (resource, params) => {
        const url = `${apiUrl}/${resource}`;
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
        const url = `${apiUrl}/${resource}`;

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
    //     return httpClient(`${apiUrl}/${resource}?${stringify(query)}`, {
    //         method: 'DELETE',
    //         body: JSON.stringify(params.data),
    //     }).then(({ json }) => ({ data: json }));
    // },

    // getOne: (resource, params) =>
    //     httpClient(`${apiUrl}/${resource}/${params.id}`)
    //     .then(({ json }) => ({
    //         data: json,
    //     })),

    // getMany: (resource, params) => {
    //     const query = {
    //         filter: JSON.stringify({ id: params.ids }),
    //     };
    //     const url = `${apiUrl}/${resource}?${stringify(query)}`;
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
    //     const url = `${apiUrl}/${resource}?${stringify(query)}`;

    //     return httpClient(url).then(({ headers, json }) => ({
    //         data: json,
    //         total: parseInt(headers.get('content-range').split('/').pop(), 10),
    //     }));
    // },

    // update: (resource, params) =>
    //     httpClient(`${apiUrl}/${resource}/${params.id}`, {
    //         method: 'PUT',
    //         body: JSON.stringify(params.data),
    //     }).then(({ json }) => ({ data: json })),

    // updateMany: (resource, params) => {
    //     const query = {
    //         filter: JSON.stringify({ id: params.ids }),
    //     };
    //     return httpClient(`${apiUrl}/${resource}?${stringify(query)}`, {
    //         method: 'PUT',
    //         body: JSON.stringify(params.data),
    //     }).then(({ json }) => ({ data: json }));
    // },


};