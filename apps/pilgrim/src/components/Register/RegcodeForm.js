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

const CheckRegcodeForm = (props) => {

    const notify = useNotify();
    const [regcode, setRegcode] = useState('');

    const submit = (event) => {
        event.preventDefault();
        notify('Validating registration code...')
        return axios({
            method: 'POST',
            url: authProviderUrl + "/validate-regcode",
            headers: {
                'Subject': 'Pilgrim',
            },
            params: {
                regcode: regcode,
            }
        })
            .then(response => {
                if (response.status < 200 || response.status >= 300) {
                    notify(response.statusText);
                } else {
                    notify(response.data["message"]);
                    props.onValidRegcode(response.data["username"], regcode);
                }
                return;
            })
            .catch(() => notify('Registration failed.'));;
    }

    if (props.form) {
        return (
            <div className={props.classes.section}>
                <Typography component="h1" variant="body1" color="textSecondary" gutterBottom>
                    To create a user account, you need a unique registration code which can be obtained from your organization representative.
                </Typography>


                <form className={props.classes.form} onSubmit={submit} noValidate>
                    <TextField
                        variant="outlined"
                        margin="normal"
                        required
                        fullWidth
                        id="regcode"
                        label="Registration Code"
                        name="vic"
                        autoComplete="vic"
                        autoFocus
                        value={regcode}
                        onChange={e => setRegcode(e.target.value)}
                    />

                    <Button
                        type="submit"
                        fullWidth
                        variant="contained"
                        color="primary"
                        className={props.classes.submit}
                    >
                        CHECK REGISTRATION CODE
                    </Button>
                </form>
            </div>
        );
    } else {
        return null;
    }

}

export default CheckRegcodeForm;