import React, { useState, useEffect } from "react";
import Card from '@material-ui/core/Card';
import { makeStyles } from '@material-ui/core/styles';
import { Title } from 'react-admin';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import axios from 'axios';

var dataProviderUrl;
if (!process.env.NODE_ENV || process.env.NODE_ENV === 'development') {
    // dev code
    dataProviderUrl = process.env.REACT_APP_KUBEINN_SCHUTTERIJ_URL + '/api/innkeeper';
} else {
    // production code
    dataProviderUrl = '/api/innkeeper';
}
console.log(dataProviderUrl)

function getCookie(name) {
    const value = `; ${document.cookie}`;
    const parts = value.split(`; ${name}=`);
    if (parts.length === 2) return parts.pop().split(';').shift();
}

const useStyles = makeStyles({
    table: {
        minWidth: 650,
    },
});

export default function Dashboard() {
    const classes = useStyles();

    const [data, setData] = useState([])

    useEffect(() => {
        axios({
            url: dataProviderUrl + '/pods',
            method: 'GET',
            headers: {
                'Authorization': getCookie("Authorization"),
            },
            timeout: 5000,
            responseType: 'json',
            responseEncoding: 'utf8',
        }).then(json => setData(json.data))
    }, [])

    console.log(data);

    if (data == null){
        return(
            <Card>No values to display</Card>
        );
    } else {
        return (
            <Card>
                <Title title="Dashboard" />
                <TableContainer component={Paper}>
                    <Table className={classes.table} size="small" aria-label="a dense table">
                        <TableHead>
                            <TableRow>
                                <TableCell align="right">pod</TableCell>
                                <TableCell align="right">namespace</TableCell>
                                <TableCell align="right">created_by_name</TableCell>
                                <TableCell align="right">node</TableCell>
                                <TableCell align="right">kube_pod_created</TableCell>
                                <TableCell align="right">kube_pod_completed</TableCell>
                                <TableCell align="right">container_cpu_usage_seconds_total</TableCell>
                                <TableCell align="right">container_memory_usage_bytes</TableCell>
                                <TableCell align="right">kube_pod_container_status_running</TableCell>
                                <TableCell align="right">kube_pod_container_status_terminated</TableCell>
                                <TableCell align="right">kube_pod_container_status_terminated_reason</TableCell>
                            </TableRow>
                        </TableHead>
                        <TableBody>
                            {data.map((row) => (
                                <TableRow key={row.pod}>
                                    <TableCell align="right">{row.pod}</TableCell>
                                    <TableCell align="right">{row.namespace}</TableCell>
                                    <TableCell align="right">{row.created_by_name}</TableCell>
                                    <TableCell align="right">{row.node}</TableCell>
                                    <TableCell align="right">{row.kube_pod_created}</TableCell>
                                    <TableCell align="right">{row.kube_pod_completed}</TableCell>
                                    <TableCell align="right">{row.container_cpu_usage_seconds_total}</TableCell>
                                    <TableCell align="right">{row.container_memory_usage_bytes}</TableCell>
                                    <TableCell align="right">{row.kube_pod_container_status_running}</TableCell>
                                    <TableCell align="right">{row.kube_pod_container_status_terminated}</TableCell>
                                    <TableCell align="right">{row.kube_pod_container_status_terminated_reason}</TableCell>
                                </TableRow>
                            ))}
                        </TableBody>
                    </Table>
                </TableContainer>
            </Card>
        );
    }
}