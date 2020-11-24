import React from 'react';
import './App.css';
import Button from '@material-ui/core/Button';
import Card from '@material-ui/core/Card';
import CardContent from '@material-ui/core/CardContent';
import Typography from '@material-ui/core/Typography';
import ButtonGroup from '@material-ui/core/ButtonGroup';
import Container from '@material-ui/core/Container';
import Particles from 'react-particles-js';
import particlesConfig from './config/particlesConfig';
import { makeStyles } from '@material-ui/core/styles';

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

function App() {
  const classes = useStyles();

  return (
    <div className="App">
      <div style={{ position: 'absolute'}}>
        <Particles height="100vh" width="100vw" params={particlesConfig} />
      </div>
      <header className="App-header">
        <Container component="main" maxWidth="xs">
          <div className={classes.paper}>
            <Card className={classes.root}>
              <CardContent>
                <div className={classes.section}>
                  <Typography component="h1" variant="h5">Welcome to Kubeinn!</Typography>
                </div>
                <div className={classes.section}>
                  <Typography component="h1" variant="body1" color="textSecondary" gutterBottom>
                    To begin, please tell me who you are:
                  </Typography>
                </div>
                <div className={classes.section}>
                  <ButtonGroup
                    orientation="vertical"
                    color="primary"
                    aria-label="vertical outlined primary button group"
                  >
                    <Button href="/innkeeper/">Cluster Administrator</Button>
                    <Button href="/reeve/">Project Manager</Button>
                    <Button href="/pilgrim/">Project User</Button>
                  </ButtonGroup>
                </div>
              </CardContent>
            </Card>
          </div>
        </Container>
      </header>
    </div>
  );
}

export default App;