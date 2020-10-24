import React, { useState } from 'react';
import { useLogin, useNotify, Notification } from 'react-admin';
import Avatar from '@material-ui/core/Avatar';
import Button from '@material-ui/core/Button';
import ButtonGroup from '@material-ui/core/ButtonGroup';
import CssBaseline from '@material-ui/core/CssBaseline';
import TextField from '@material-ui/core/TextField';
import FormControlLabel from '@material-ui/core/FormControlLabel';
import Checkbox from '@material-ui/core/Checkbox';
import Card from '@material-ui/core/Card';
import CardContent from '@material-ui/core/CardContent';
import Divider from '@material-ui/core/Divider';
import Link from '@material-ui/core/Link';
import Grid from '@material-ui/core/Grid';
import Box from '@material-ui/core/Box';
import Radio from '@material-ui/core/Radio';
import RadioGroup from '@material-ui/core/RadioGroup';
import Typography from '@material-ui/core/Typography';
import { makeStyles } from '@material-ui/core/styles';
import RegisterVillageForm from './RegisterVillageForm';
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
        margin: theme.spacing(4, 2),
    }
}));

const RegisterPage = () => {
    const classes = useStyles();

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
                            <Typography component="h1" variant="body1" color="textSecondary" gutterBottom>
                                To use this platform, organization representatives must first request for a Village Identification Code (VIC).
                            </Typography>
                            <Typography component="h1" variant="body1" color="textSecondary" gutterBottom>
                                By submitting this form, I understand that I am responsible for all users registered under my VIC.
                            </Typography>
                        </div>
                        <Divider variant="middle" />
                        <RegisterVillageForm classes={classes} />
                        <div className={classes.section}>
                            <Grid container justify="flex-end">
                                <Grid item>
                                    <Link href="login" variant="body2">Already have an account? Sign in</Link>
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

export default RegisterPage;