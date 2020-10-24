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
        margin: theme.spacing(4, 2),
    }
}));

const RegisterPage = () => {
    const classes = useStyles();
    const [form, setForm] = useState(0);
    const handleChange = (event) => {
        setForm(event.target.value);
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

                        <div className={classes.section}>
                            <Typography component="h1" variant="body1" color="textSecondary" gutterBottom>
                                A Village Identification Code (VIC) is required for registration.
                            </Typography>
                            <RadioGroup aria-label="VIC" name="vic" value={form} onChange={handleChange}>
                                <FormControlLabel value="true" control={<Radio />} label="I know my VIC." />
                                <FormControlLabel value="false" control={<Radio />} label="I do not know my VIC." />
                            </RadioGroup>
                        </div>
                        <Divider variant="middle" />
                        <DisplayForm classes={classes} displayForm={form} />
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

function DisplayForm(props) {
    const displayForm = props.displayForm;
    if (displayForm === "true") {
        return <RegisterPilgrimForm classes={props.classes} />;
    } else if (displayForm === "false") {
        return (
            <div className={props.classes.section}>
                <Typography component="h1" variant="body1" color="textSecondary" gutterBottom>
                    Please contact your organization representative to obtain your VIC.
                </Typography>
            </div>
        );
    } else {
        return null;
    }
}

export default RegisterPage;