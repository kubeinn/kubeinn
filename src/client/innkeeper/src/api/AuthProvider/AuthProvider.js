import axios from 'axios';

var authProviderUrl;
if (!process.env.NODE_ENV || process.env.NODE_ENV === 'development') {
    // dev code
    authProviderUrl = process.env.REACT_APP_KUBEINN_SCHUTTERIJ_URL + '/api/auth';
} else {
    // production code
    authProviderUrl = '/api/auth';
}
console.log(authProviderUrl)

export default {
    // called when the user attempts to log in
    login: ({ username, password }) => {
        return axios.post(authProviderUrl + "/login", {
            username: username,
            password: password,
        }, { headers: { 'Subject': 'innkeeper' } })
            .then(response => {
                if (response.status < 200 || response.status >= 300) {
                    throw new Error(response.statusText);
                }
                setCookie("Authorization", "Bearer " + response.data.Authorization, 1)
                setCookie("path", "/", 1)
                return;
            })
    },

    // called when the user clicks on the logout button
    logout: () => {
        eraseCookie("Authorization")
        return Promise.resolve();
    },
    // called when the API returns an error
    checkError: ({ status }) => {
        if (status === 401 || status === 403) {
            eraseCookie("Authorization: Bearer ")
            return Promise.reject();
        }
        return Promise.resolve();
    },
    // called when the user navigates to a new location, to check for authentication
    checkAuth: () => {
        return axios.post(authProviderUrl + "/check-auth", {
            timeout: 5000,
            responseType: 'json',
            responseEncoding: 'utf8',
        }, {
            headers: {
                'Subject': 'Innkeeper',
                'Authorization': getCookie("Authorization")
            }
        })
            .then(response => {
                if (response.status < 200 || response.status >= 300) {

                    return Promise.reject()
                }
                else {
                    return Promise.resolve()
                }

            })
    },
    // called when the user navigates to a new location, to check for permissions / roles
    getPermissions: () => Promise.resolve(),
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