import * as React from "react";
import {
    List,
    Edit,
    Create,
    Datagrid,
    TextField,
    EmailField,
    SimpleForm,
    TextInput,
    EditButton,
} from 'react-admin';

export const InnkeeperList = props => (
    <List {...props} >
        <Datagrid>
            <TextField source="id" label="InnkeeperID" />
            <TextField source="username" label="Username" />
            <EmailField source="email" label="Email" />
        </Datagrid>
    </List>
);

// export const InnkeeperCreate = props => (
//     <Create {...props}>
//         <SimpleForm>
//             <TextInput source="username" label="Username" />
//             <TextInput source="email" label="Email" />
//             <PasswordInput label="Password" source="passwd" helperText="Default password: innkeeper" initiallyVisible/>
//         </SimpleForm>
//     </Create>
// );

// export const InnkeeperEdit = (props) => {
//     return (
//         <Edit {...props} >
//             <SimpleForm>
//                 <TextInput source="email" label="Email" />
//             </SimpleForm>
//         </Edit>
//     );
// }