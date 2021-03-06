import React, { useState } from 'react';
import { Notification } from 'react-admin';
import CssBaseline from '@material-ui/core/CssBaseline';
import Card from '@material-ui/core/Card';
import CardContent from '@material-ui/core/CardContent';
import Divider from '@material-ui/core/Divider';
import { Link } from 'react-router-dom'
import Grid from '@material-ui/core/Grid';
import Typography from '@material-ui/core/Typography';
import { makeStyles } from '@material-ui/core/styles';
import RegisterPilgrimForm from './RegisterPilgrimForm';
import Container from '@material-ui/core/Container';

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
    }
}));

const RegisterPage = () => {
    const classes = useStyles();

    const [registrationNotice, setRegistrationNotice] = useState(false);
    const [form, setForm] = useState(true);

    const handleChange = () => {
        setForm(false);
        setRegistrationNotice(true);
    }

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
                        <div className={classes.section}>
                            <RegisterPilgrimForm classes={classes} form={form} onSuccessfulRegistration={handleChange} />
                            <DisplayRegistrationNotice classes={classes} registrationNotice={registrationNotice} />
                        </div>
                        <div className={classes.section}>
                            <Grid container justify="flex-end">
                                <Grid item>
                                    <Link to="/login" variant="body2">Already have an account? Sign in</Link>
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

function DisplayRegistrationNotice(props) {
    const registrationNotice = props.registrationNotice;
    if (registrationNotice) {
        return (
            <Typography component="h1" variant="body1" color="textSecondary" gutterBottom>
                Your request has been submitted. Our administrators will review your request and follow up via the email provided.
            </Typography>
        );
    } else {
        return null;
    }
}

export default RegisterPage;