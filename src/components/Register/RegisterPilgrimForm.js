import React from 'react';
import { useState } from 'react';
import { withStyles } from '@material-ui/core/styles';
import Button from '@material-ui/core/Button';
import TextField from '@material-ui/core/TextField';
import Link from '@material-ui/core/Link';
import Grid from '@material-ui/core/Grid';
import { useLogin, useNotify, Notification } from 'react-admin';
import axios from 'axios';

// Production
// const authProviderUrl = window._env_.KUBEINN_SCHUTTERIJ_URL + '/auth';
// Local
const authProviderUrl = process.env.REACT_APP_KUBEINN_SCHUTTERIJ_URL + '/auth';

const RegisterPilgrimForm = (props) => {
    const notify = useNotify();

    const [vic, setOid] = useState('');
    const [username, setUsername] = useState('');
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');

    const submit = (event) => {
        event.preventDefault();
        notify('Registering new user...')
        return axios({
            method: 'POST',
            url: authProviderUrl + "/register-pilgrim",
            headers: {
                'Subject': 'Pilgrim',
            },
            params: {
                vic: vic,
                username: username,
                email: email,
                password: password,
            }
        })
            .then(response => {
                if (response.status < 200 || response.status >= 300) {
                    notify(response.statusText);
                } else {
                    notify(response.data["Message"]);
                }
                return;
            })
            .catch(() => notify('Registration failed. Please contact administrator.'));;
    }


    return (
        <form className={props.classes.form} onSubmit={submit} noValidate>
            <TextField
                variant="outlined"
                margin="normal"
                required
                fullWidth
                id="vic"
                label="Village Identification Code (VIC)"
                name="vic"
                autoComplete="vic"
                autoFocus
                value={vic}
                onChange={e => setOid(e.target.value)}
            />
            <TextField
                variant="outlined"
                margin="normal"
                required
                fullWidth
                id="username"
                label="Username"
                name="username"
                autoComplete="username"
                value={username}
                onChange={e => setUsername(e.target.value)}
            />
            <TextField
                variant="outlined"
                margin="normal"
                required
                fullWidth
                id="email"
                label="Email"
                name="email"
                value={email}
                onChange={e => setEmail(e.target.value)}
            />
            <TextField
                variant="outlined"
                margin="normal"
                required
                fullWidth
                name="password"
                label="Password"
                type="password"
                id="password"
                autoComplete="current-password"
                value={password}
                onChange={e => setPassword(e.target.value)}
            />
            <Button
                type="submit"
                fullWidth
                variant="contained"
                color="primary"
                className={props.classes.submit}
            >
                Register
                    </Button>
        </form>
    );
}

export default RegisterPilgrimForm;