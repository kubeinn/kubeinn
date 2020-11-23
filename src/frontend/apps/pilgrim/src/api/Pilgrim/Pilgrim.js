import axios from 'axios';

var pilgrimAPI;
if (!process.env.NODE_ENV || process.env.NODE_ENV === 'development') {
    // dev code
    pilgrimAPI = process.env.REACT_APP_KUBEINN_SCHUTTERIJ_URL + '/pilgrim';
} else {
    // production code
    pilgrimAPI = window._env_.KUBEINN_SCHUTTERIJ_URL + '/pilgrim';
}
console.log(pilgrimAPI)


export async function CreateProjectByPilgrim(params) {
    console.log("Creating project...");

    const url = `${pilgrimAPI}/create-project`;

    var response = await axios({
        url: url,
        method: 'POST',
        headers: {
            'Authorization': getCookie("Authorization"),
        },
        params: {
            namespace: params.data.title,
            cpu: params.data.cpu,
            memory: params.data.memory,
            storage: params.data.storage,
        },
        timeout: 60000,
        responseType: 'json',
        responseEncoding: 'utf8',
    });

    return response;
}

export function DeleteProjectByPilgrim(params) {
    console.log("Deleting project...");

    const url = `${pilgrimAPI}/delete-project`;

    return axios({
        url: url,
        method: 'POST',
        headers: {
            'Authorization': getCookie("Authorization"),
        },
        params: {
            id: params.data.id,
        },
        timeout: 60000,
        responseType: 'json',
        responseEncoding: 'utf8',
    });
}

export async function DeleteProjectsByPilgrim(params) {
    console.log("Deleting projects...");

    const url = `${pilgrimAPI}/delete-project`;
    var arrayLength = params.ids.length;
    var promises = []; // store all promise

    for (let i = 0; i < arrayLength; i++) {
        promises.push(
            axios({
                url: url,
                method: 'POST',
                headers: {
                    'Authorization': getCookie("Authorization"),
                },
                params: {
                    id: params.ids[i],
                },
                timeout: 60000,
                responseType: 'json',
                responseEncoding: 'utf8',
            }));
    }
    await Promise.all(promises);
    return;
}

function getCookie(name) {
    const value = `; ${document.cookie}`;
    const parts = value.split(`; ${name}=`);
    if (parts.length === 2) return parts.pop().split(';').shift();
}