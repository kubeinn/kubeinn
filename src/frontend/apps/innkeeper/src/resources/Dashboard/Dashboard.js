import * as React from "react";
import Card from '@material-ui/core/Card';
import CardContent from '@material-ui/core/CardContent';
import { Title } from 'react-admin';

// Production
// const grafanaUrl = window._env_.KUBEINN_GRAFANA_URL
// Local
// const grafanaUrl = process.env.REACT_APP_KUBEINN_GRAFANA_URL;

class Dashboard extends React.Component {
    render() {
        return (
            <div>
                <Card>
                    <Title title="Dashboard" />
                    <CardContent min-height="100vh">
                        <iframe title="grafana" src={grafanaUrl} style={{width: "100%", height: "100vh"}}/>
                    </CardContent>
                </Card >
            </div>
        );
    }
}

export default Dashboard;