import * as React from 'react';
import { useState } from 'react';
import { useLogin, useNotify, Notification } from 'react-admin';
import Button from '@material-ui/core/Button';
import CssBaseline from '@material-ui/core/CssBaseline';
import TextField from '@material-ui/core/TextField';
import Card from '@material-ui/core/Card';
import CardContent from '@material-ui/core/CardContent';
import FormControlLabel from '@material-ui/core/FormControlLabel';
import Checkbox from '@material-ui/core/Checkbox';
import Link from '@material-ui/core/Link';
import Grid from '@material-ui/core/Grid';
import Box from '@material-ui/core/Box';
import LockOutlinedIcon from '@material-ui/icons/LockOutlined';
import Typography from '@material-ui/core/Typography';
import Divider from '@material-ui/core/Divider';
import Container from '@material-ui/core/Container';
import { makeStyles } from '@material-ui/core/styles';

// Production
// const reeveUrl = window._env_.KUBEINN_POSTGREST_URL;
// Local
const reeveUrl = process.env.REACT_APP_KUBEINN_REEVE_URL;

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
        marginTop: theme.spacing(1),
    },
    submit: {
        margin: theme.spacing(3, 0, 2),
    },
    section: {
        margin: theme.spacing(3, 2),
    },
}));

const LoginForm = (props) => {
    const notify = useNotify();
    const login = useLogin();

    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');

    const submit = (e) => {
        e.preventDefault();
        login({ username, password })
            .catch(() => notify('Invalid username or password'));
    };

    return (
        <form className={props.classes.form} onSubmit={submit} noValidate>
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
                LOG IN
                </Button>
        </form>
    );
}

const LoginPage = () => {
    const classes = useStyles();

    return (
        <Container component="main" maxWidth="xs">
            <CssBaseline />
            <div className={classes.paper}>
                <Card className={classes.root}>
                    <CardContent>
                        <div className={classes.section}>
                            <Typography component="h1" variant="h5">Pilgrim User Portal</Typography>
                        </div>
                        <Divider variant="middle" />
                        <div className={classes.section}>
                            <LoginForm classes={classes} />
                        </div>
                        <div className={classes.section}>
                            <Grid container justify="flex-end">
                                <Grid item>
                                    <Link href="register" variant="body2">
                                        {"Don't have an account? Register"}
                                    </Link>
                                </Grid>
                            </Grid>
                        </div>
                        <Notification />
                    </CardContent>
                </Card>
            </div>
        </Container>

    );
};

export default LoginPage;