import * as React from "react";
import Card from '@material-ui/core/Card';
import CardContent from '@material-ui/core/CardContent';
import { Title } from 'react-admin';
// import { Dashboard } from "../../Resources/Dashboard/Dashboard";

class Dashboard extends React.Component {

    render() {
        return (
            <div>
                <Card>
                    <Title title="Pilgrim User Portal" />
                    <CardContent>Lorem ipsum sic dolor amet...</CardContent>
                </Card >
            </div>
        );
    }
}

export default Dashboard;