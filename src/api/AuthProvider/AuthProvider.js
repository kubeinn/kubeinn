import axios from 'axios';

// Production
// const authProviderUrl = window._env_.KUBEINN_SCHUTTERIJ_URL + '/auth';
// Local
const authProviderUrl = process.env.REACT_APP_KUBEINN_SCHUTTERIJ_URL + '/auth';

export default {
    // called when the user attempts to log in
    login: ({ username, password }) => {
        return axios.post(authProviderUrl + "/validate-user", {
            params: {
                username: username,
                password: password,
            },
            timeout: 5000,
            responseType: 'json',
            responseEncoding: 'utf8',
        }, { headers: { 'Subject': 'Pilgrim' } })
            .then(response => {
                if (response.status < 200 || response.status >= 300) {
                    throw new Error(response.statusText);
                }
                return response.json();
            })
            .then(({ data }) => {
                console.log(data);
                document.cookie = "Authorization: Bearer " + data;
            });
    },

    // called when the user clicks on the logout button
    logout: () => {
        localStorage.removeItem('username');
        return Promise.resolve();
    },
    // called when the API returns an error
    checkError: ({ status }) => {
        if (status === 401 || status === 403) {
            localStorage.removeItem('username');
            return Promise.reject();
        }
        return Promise.resolve();
    },
    // called when the user navigates to a new location, to check for authentication
    checkAuth: () => {
        return localStorage.getItem('username')
            ? Promise.resolve()
            : Promise.reject();
    },
    // called when the user navigates to a new location, to check for permissions / roles
    getPermissions: () => Promise.resolve(),
};