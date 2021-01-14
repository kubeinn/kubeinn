import React from 'react';
import { useState } from 'react';
import Button from '@material-ui/core/Button';
import TextField from '@material-ui/core/TextField';
import Checkbox from '@material-ui/core/Checkbox';
import FormControlLabel from '@material-ui/core/FormControlLabel';
import { useNotify } from 'react-admin';
import axios from 'axios';

var authProviderUrl;
if (!process.env.NODE_ENV || process.env.NODE_ENV === 'development') {
    // dev code
    authProviderUrl = process.env.REACT_APP_KUBEINN_SCHUTTERIJ_URL + '/auth';
} else {
    // production code
    authProviderUrl = '/api/auth';
}
console.log(authProviderUrl)

const RegisterVillageForm = (props) => {
    const notify = useNotify();
    const form = props.form;

    const [organization, setOrganization] = useState('');
    const [description, setDescription] = useState('');
    const [username, setUsername] = useState('');
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [checkedTerms, setCheckedTerms] = useState(false);

    const handleChangeCheckbox = (event) => {
        setCheckedTerms(event.target.checked);
    };

    const submit = (event) => {
        event.preventDefault();

        if (!checkedTerms) {
            notify('Registration failed. You must agree to the T&Cs above.')
            return;
        }

        notify('Registering new pilgrim...')
        return axios({
            method: 'POST',
            url: authProviderUrl + "/register/pilgrim",
            headers: {

            },
            params: {
                organization: organization,
                description: description,
                username: username,
                email: email,
                password: password,
            }
        })
            .then(response => {
                if (response.status < 200 || response.status >= 300) {
                    notify(response.statusText);
                } else {
                    notify(response.data["message"]);
                    props.onSuccessfulRegistration();
                }
                return;
            })
            .catch(() => notify('Registration failed. Please contact administrator.'));;
    }

    if (form) {
        return (
            <div className={props.classes.section}>
                <form className={props.classes.form} onSubmit={submit} noValidate>
                    <TextField
                        variant="outlined"
                        margin="normal"
                        required
                        fullWidth
                        id="organization"
                        label="Organization"
                        name="organization"
                        autoComplete="organization"
                        value={organization}
                        onChange={e => setOrganization(e.target.value)}
                    />
                    <TextField
                        variant="outlined"
                        margin="normal"
                        required
                        fullWidth
                        id="description"
                        label="Description of Organization"
                        name="description"
                        value={description}
                        onChange={e => setDescription(e.target.value)}
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
                    <FormControlLabel
                        control={<Checkbox checked={checkedTerms} onChange={handleChangeCheckbox} name="checkbox" />}
                        label="I understand that I will be held responsible for all users registered under my organization."
                    />
                    <Button
                        type="submit"
                        fullWidth
                        variant="contained"
                        color="primary"
                        className={props.classes.submit}
                    >
                        Submit
                    </Button>
                </form>
            </div>
        );
    } else {
        return null;
    }
}

export default RegisterVillageForm;