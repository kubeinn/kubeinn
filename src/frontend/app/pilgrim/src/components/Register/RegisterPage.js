import React, { useState } from 'react';
import { Notification } from 'react-admin';
import Button from '@material-ui/core/Button';
import CssBaseline from '@material-ui/core/CssBaseline';
import Card from '@material-ui/core/Card';
import CardContent from '@material-ui/core/CardContent';
import Divider from '@material-ui/core/Divider';
import Link from '@material-ui/core/Link';
import Grid from '@material-ui/core/Grid';
import Typography from '@material-ui/core/Typography';
import { makeStyles } from '@material-ui/core/styles';
import PasswordForm from './PasswordForm';
import RegcodeForm from './RegcodeForm';
import Container from '@material-ui/core/Container';

var reeveUrl;
if (!process.env.NODE_ENV || process.env.NODE_ENV === 'development') {
    // dev code
    reeveUrl = process.env.REACT_APP_KUBEINN_REEVE_URL;
} else {
    // production code
    reeveUrl = '/reeve/';
}
console.log(reeveUrl)

const useStyles = makeStyles(theme => ({
    paper: {
        display: 'flex',
        flexDirection: 'column',
        alignItems: 'center',
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
    card: {
        margin: theme.spacing(4, 0),
    },
}));

const RegisterPage = () => {
    const classes = useStyles();

    const [regcodeForm, setRegcodeForm] = useState(true);
    const [passwordForm, setPasswordForm] = useState(false);
    const [username, setUsername] = useState("default");
    const [regcode, setRegcode] = useState("");

    const onValidRegcode = (username, regcode) => {
        setPasswordForm(true);
        setRegcodeForm(false);
        setUsername(username);
        setRegcode(regcode);
    };

    const onPasswordSet = () => {
        setPasswordForm(false);
    };

    return (
        <Container component="main" maxWidth="xs" >
            <CssBaseline />
            <div className={classes.paper}>
                <Card className={classes.card}>
                    <CardContent>
                        <div className={classes.section}>
                            <Typography component="h1" variant="h5">Registration</Typography>
                        </div>
                        <Divider variant="middle" />
                        <RegcodeForm classes={classes} form={regcodeForm} onValidRegcode={onValidRegcode} />
                        <PasswordForm classes={classes} form={passwordForm} username={username} regcode={regcode} onPasswordSet={onPasswordSet} />
                        <RegistrationNotice classes={classes} regcodeForm={regcodeForm} passwordForm={passwordForm} />
                        <Divider variant="middle" />
                        <div className={classes.section}>
                            <Typography component="h1" variant="body1" color="textSecondary" gutterBottom>
                                If you are the representative for your organization, click <Link href={reeveUrl} variant="body2">here</Link>.
                            </Typography>
                        </div>
                        <Divider variant="middle" />
                        <div className={classes.section}>
                            <Grid container justify="flex-end">
                                <Grid item>
                                    <Link href="/pilgrim/login" variant="body2">Already have an account? Sign in</Link>
                                </Grid>
                            </Grid>
                        </div>
                    </CardContent>
                </Card>
                <Notification />
            </div>
        </Container>
    );
}

const RegistrationNotice = (props) => {
    if (!props.regcodeForm && !props.passwordForm) {
        return (
            <div className={props.classes.section}>
                <Button variant="contained" fullWidth color="primary" href="/pilgrim/login">
                    PROCEED TO SIGN IN.
                </Button>
            </div>
        );
    } else {
        return null;
    }
}

export default RegisterPage;