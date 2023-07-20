
import PaginationExampleWithPageSizeSelector from "@/components/DataTableWrapper";
import CreateOrUpdateData from "@/content/users/CreateOrUodateData";
import { useDeleteUserMutation, useGetAllUsersQuery } from "@/redux/services/user";
import { User } from "@/types/users";
import { ActionIcon, Group, Header, Text } from "@mantine/core";
import { IconEdit, IconTrash } from "@tabler/icons-react";
import { useState } from "react";


export default function UserIndex() {
  // const { users: any } = useLoaderData();

  const {data:users,refetch}= useGetAllUsersQuery();
  console.log({users})
  const [deleteUser]= useDeleteUserMutation();
  // const {data:me,isLoading}= useAuthUserQuery();

  
  const [editUser, setEditUser] = useState();
  console.log({ editUser })

const  handleClick = ()=>{ 
  // refetch();
}
const handleDelete = async(id:number)=>{
  const {data,error} = await deleteUser(id);
  console.log({data,error});
}
  return (
    <>
    {/* <Button bg={"red"} onClick={handleClick} variant={"filled"}> CLICK</Button> */}
      {/* <PaginateTableWrapper/> */}
      <Header height={64} display="flex" sx={{ justifyContent: 'space-between' }}>
        <Text pt={18} pl={18} size={"lg"} weight={600} color="orange">Users</Text>
        {/* <Text pt={18} pl={18} size={"lg"} weight={600} color="orange">Users</Text> */}
        <CreateOrUpdateData editData={editUser} setEditData={setEditUser} refetch={refetch} />
      </Header>
        
        {/* <Button variant="gradient" gradient={{ from: 'orange', to: 'red' }}>Orange red</Button> */}

      <div className="px-6 py-3 w-full min-h-fit">
        <PaginationExampleWithPageSizeSelector
          headColumns={[
            { accessor: 'id', width: 100 },
            { accessor: 'first_name'},
            { accessor: 'last_name'},
            { accessor: 'username'},
            { accessor: 'email'},
            { accessor: 'phone_number'},
            {
              accessor: 'actions',
              title: <Text mr="xs">Row actions</Text>,
              textAlignment: 'right',
              render: (data:User) => (
                <Group spacing={4} position="right" noWrap>
                  {/* <ActionIcon color="green" onClick={() => showInfo(data)}>
                    <IconEye size={16} />
                  </ActionIcon> */}
                  <ActionIcon color="sky" onClick={() => setEditUser(data)}>
                    <IconEdit size={20} />
                  </ActionIcon>
                  <ActionIcon color="red" onClick={() => handleDelete(data.id)}>
                    <IconTrash size={20} />
                  </ActionIcon>
                </Group>
              ),
            },
          ]}
          datas={ users?.users || []}
        />
      </div>

      {/* <Table verticalSpacing="xs" fontSize="xs">

        <thead>
          <tr>
            <td>Username</td>
          </tr>
        </thead>

        <tbody>
          {
            users?.map((user: unknown, index: number) => (<td key={index}>{user.username}</td>))
          }
        </tbody>
      </Table> */}
      {/* <Table columns={columns} dataSource={users}/> */}
    </>
  )
}


