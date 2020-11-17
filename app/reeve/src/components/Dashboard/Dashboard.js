import React, { useState } from 'react';
import Typography from '@material-ui/core/Typography';

const Dashboard = () => {
    return (
        <Typography component="h1" variant="body1" color="textSecondary" gutterBottom>
            To use this platform, organization representatives must first request for a Village Identification Code (VIC).
        </Typography>
    );
}

export default Dashboard;