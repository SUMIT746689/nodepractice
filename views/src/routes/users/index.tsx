
import PaginationExampleWithPageSizeSelector from "@/components/DataTableWrapper";
import CreateOrUpdateData from "@/content/users/CreateOrUodateData";
import { useGetAllUsersQuery } from "@/redux/services/user";
import { ActionIcon, Button, Group, Header, Text } from "@mantine/core";
import { IconEdit } from "@tabler/icons-react";
// import { IconEye } from "@tabler/icons-react";
import { useState } from "react";


export default function UserIndex() {
  // const { users: any } = useLoaderData();

  const {data:users}= useGetAllUsersQuery();
  console.log({users})
  // const {data:me,isLoading}= useAuthUserQuery();

  
  const [editUser, setEditUser] = useState();
  console.log({ editUser })

const  handleClick = ()=>{ 
  // refetch();
}
  return (
    <>
    {/* <Button bg={"red"} onClick={handleClick} variant={"filled"}> CLICK</Button> */}
      {/* <PaginateTableWrapper/> */}
      <Header height={64} display="flex" sx={{ justifyContent: 'space-between' }}>
        <Text pt={18} pl={18} size={"lg"} weight={600} color="orange">Users</Text>
        {/* <Text pt={18} pl={18} size={"lg"} weight={600} color="orange">Users</Text> */}
        <CreateOrUpdateData editData={editUser} setEditData={setEditUser} />
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
              render: (data) => (
                <Group spacing={4} position="right" noWrap>
                  {/* <ActionIcon color="green" onClick={() => showInfo(data)}>
                    <IconEye size={16} />
                  </ActionIcon> */}
                  <ActionIcon color="orange" onClick={() => setEditUser(data)}>
                    <IconEdit size={20} />
                  </ActionIcon>
                  {/* <ActionIcon color="red" onClick={() => deleteCompany(data)}>
                    <IconTrash size={16} />
                  </ActionIcon> */}
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


