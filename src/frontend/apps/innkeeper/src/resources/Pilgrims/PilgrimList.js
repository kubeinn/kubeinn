import * as React from "react";
import {
    List,
    Datagrid,
    TextField,
    EmailField,
} from 'react-admin';

export const PilgrimList = props => (
    <List {...props} >
        <Datagrid>
            <TextField source="id" label="PilgrimID" />
            <TextField source="villageid" label="VillageID" />
            <TextField source="username" label="Username" />
            <EmailField source="email" label="Email" />
            <TextField source="regcode" label="Registration Code" />
        </Datagrid>
    </List>
);

// export const PilgrimCreate = props => (
//     <Create {...props}>
//         <SimpleForm>
//             <TextInput source="villageid" label="VillageID" />
//             <TextInput source="username" label="Username" />
//             <TextInput source="email" label="Email" />
//             <TextInput source="villageid" label="VillageID" />
//         </SimpleForm>
//     </Create>
// );

// export const PilgrimEdit = (props) => {
//     return (
//         <Edit {...props} >
//             <SimpleForm>
//                 <TextInput source="email" label="Email" />
//                 <TextInput source="regcode" label="Registration Code" />
//             </SimpleForm>
//         </Edit>
//     );
// }