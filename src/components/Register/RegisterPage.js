import React from 'react';
import { useState } from 'react';
import Avatar from '@material-ui/core/Avatar';
import Button from '@material-ui/core/Button';
import CssBaseline from '@material-ui/core/CssBaseline';
import TextField from '@material-ui/core/TextField';
import FormControlLabel from '@material-ui/core/FormControlLabel';
import Checkbox from '@material-ui/core/Checkbox';
import Link from '@material-ui/core/Link';
import Grid from '@material-ui/core/Grid';
import Box from '@material-ui/core/Box';
import LockOutlinedIcon from '@material-ui/icons/LockOutlined';
import Typography from '@material-ui/core/Typography';
import { makeStyles } from '@material-ui/core/styles';
import Container from '@material-ui/core/Container';
import { useLogin, useNotify, Notification } from 'react-admin';
import axios from 'axios';

// Production
// const authProviderUrl = window._env_.KUBEINN_SCHUTTERIJ_URL + '/auth';
// Local
const authProviderUrl = process.env.REACT_APP_KUBEINN_SCHUTTERIJ_URL + '/auth';

const useStyles = makeStyles((theme) => ({
    paper: {
        marginTop: theme.spacing(8),
        display: 'flex',
        flexDirection: 'column',
        alignItems: 'center',
    },
    avatar: {
        margin: theme.spacing(1),
        backgroundColor: theme.palette.secondary.main,
    },
    form: {
        width: '100%', // Fix IE 11 issue.
        marginTop: theme.spacing(3),
    },
    submit: {
        margin: theme.spacing(3, 0, 2),
    },
}));

const RegisterPage = () => {
    const classes = useStyles();
    const notify = useNotify();
    const [username, setUsername] = useState('');
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const submit = (e) => {
        e.preventDefault();
        notify('Registering new user...')
        return axios({
            method: 'POST',
            url: authProviderUrl + "/register",
            headers: {
                'Subject': 'Pilgrim',
            },
            params: {
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
                    this.props.history.push('/login');
                }
                return
            })
            .catch(() => notify('Registration failed. Please contact administrator.'));;
    }

    return (
        <Container component="main" maxWidth="xs">
            <CssBaseline />
            <div className={classes.paper}>
                <Typography component="h1" variant="h5">Register</Typography>
                <form className={classes.form} onSubmit={submit} noValidate>
                    <TextField
                        variant="outlined"
                        margin="normal"
                        required
                        fullWidth
                        id="username"
                        label="Username"
                        name="username"
                        autoComplete="username"
                        autoFocus
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
                        autoFocus
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
                        className={classes.submit}
                    >
                        Register
                        </Button>
                    <Grid container justify="flex-end">
                        <Grid item>
                            <Link href="login" variant="body2">
                                Already have an account? Sign in
              </Link>
                        </Grid>
                    </Grid>
                </form>
                <Notification />
            </div>
        </Container>
    );
}

export default RegisterPage;