import React from 'react';
import { useState } from 'react';
import { withStyles } from '@material-ui/core/styles';
import Button from '@material-ui/core/Button';
import TextField from '@material-ui/core/TextField';
import Typography from '@material-ui/core/Typography';
import Link from '@material-ui/core/Link';
import Grid from '@material-ui/core/Grid';
import Card from '@material-ui/core/Card';
import CardContent from '@material-ui/core/CardContent';
import { useLogin, useNotify, Notification } from 'react-admin';
import axios from 'axios';

// Production
// const authProviderUrl = window._env_.KUBEINN_SCHUTTERIJ_URL + '/auth';
// Local
const authProviderUrl = process.env.REACT_APP_KUBEINN_SCHUTTERIJ_URL + '/auth';

const SetPasswordForm = (props) => {
    const notify = useNotify();

    const [password, setPassword] = useState('');

    const submit = (event) => {
        event.preventDefault();

        notify('Registering user...')

        return axios({
            method: 'POST',
            url: authProviderUrl + "/register/pilgrim",
            headers: {
                'Subject': 'Pilgrim',
            },
            params: {
                regcode: props.regcode,
                password: password,
            }
        })
            .then(response => {
                if (response.status < 200 || response.status >= 300) {
                    notify(response.statusText);
                } else {
                    notify(response.data["message"]);
                    props.onPasswordSet();
                }
                return;
            })
            .catch(() => notify('Registration failed.'));;
    }

    if (props.form) {
        return (
            <div className={props.classes.section}>
                <Typography component="h1" variant="body1" color="textSecondary" gutterBottom>
                    Welcome, <b>{props.username}</b>.
                </Typography>
                <form className={props.classes.form} onSubmit={submit} noValidate>
                    <TextField
                        variant="outlined"
                        margin="normal"
                        required
                        fullWidth
                        id="password"
                        label="Password"
                        name="password"
                        autoComplete="password"
                        autoFocus
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
                        SUBMIT
                    </Button>
                </form>
            </div>
        );
    } else {
        return null;
    }
}

export default SetPasswordForm;