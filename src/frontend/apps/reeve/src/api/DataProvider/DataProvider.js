import axios from 'axios';

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

    create: (resource, params) => {
        const url = `${dataProviderUrl}/${resource}`;

        return axios({
            url: url,
            method: 'POST',
            headers: {
                'Authorization': getCookie("Authorization"),
            },
            data: params.data,
            timeout: 5000,
            responseType: 'json',
            responseEncoding: 'utf8',
        }).then(({ json }) => ({
            data: { ...params.data },
        }));
    },

    delete: (resource, params) => {
        const url = `${dataProviderUrl}/${resource}`;
        const paramsId = "eq." + params.id;

        return axios({
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

        return axios({
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