
import PaginationExampleWithPageSizeSelector from "@/components/DataTableWrapper";
import CreateOrUpdateData from "@/content/users/CreateOrUodateData";
import { useAuthUserQuery } from "@/redux/services/auth";
import { isErrorWithMessage, isFetchBaseQueryError } from "@/redux/services/helpers";
import { useDeleteUserMutation, useGetAllUsersQuery } from "@/redux/services/user";
import { User } from "@/types/users";
import { ActionIcon, Group, Header, Text } from "@mantine/core";
import { notifications } from "@mantine/notifications";
import { IconEdit, IconTrash } from "@tabler/icons-react";
import { useState } from "react";


export default function UserIndex() {
  // const { users: any } = useLoaderData();

  const { data: users } = useGetAllUsersQuery();
  const { data: authUser } = useAuthUserQuery();
  console.log({ users })
  const [deleteUser] = useDeleteUserMutation();
  // const {data:me,isLoading}= useAuthUserQuery();


  const [editUser, setEditUser] = useState<User>();
  console.log({ editUser })

  // const  handleClick = ()=>{ 
  // refetch();
  // }

  const handleDelete = async (id: number): Promise<void> => {

    try {
      const { data = undefined }: { data: User | undefined } = await deleteUser(id).unwrap()
      if (data) notifications.show({ message: "sucessfully deleted" })
    } catch (err) {
      if (isFetchBaseQueryError(err)) {
        // you can access all properties of `FetchBaseQueryError` here
        const errMsg = 'error' in err ? err.error : JSON.stringify(err.data)
        notifications.show({ message: errMsg, color: 'red' })
      } else if (isErrorWithMessage(err)) {
        // you can access a string 'message' property here
        notifications.show({ message: err.message, color: 'red' })
      }
    }
    // deleteUser(id).then((data) => {
    //   console.log({ data })
    //   notifications.show({ message: "User deleted Successfully" })
    // }).catch((err: 'string') => {
    //   console.log({ err });
    //   notifications.show({ message: "Failed to deleted User",color:'red' })
    // });
  }

  return (
    <>
      {/* <Button bg={"red"} onClick={handleClick} variant={"filled"}> CLICK</Button> */}
      {/* <PaginateTableWrapper/> */}
      <Header height={64} display="flex" sx={{ justifyContent: 'space-between' }}>
        <Text pt={18} pl={18} size={"lg"} weight={600} color="orange">Users</Text>
        {/* <Text pt={18} pl={18} size={"lg"} weight={600} color="orange">Users</Text> */}
        <CreateOrUpdateData editData={editUser} setEditData={setEditUser} authUser={authUser} />
      </Header>

      {/* <Button variant="gradient" gradient={{ from: 'orange', to: 'red' }}>Orange red</Button> */}

      <div className="px-6 py-3 w-full min-h-fit">
        <PaginationExampleWithPageSizeSelector
          headColumns={[
            { accessor: 'id', width: 100 },
            { accessor: 'first_name' },
            { accessor: 'last_name' },
            { accessor: 'username' },
            { accessor: 'email' },
            { accessor: 'phone_number' },
            {
              accessor: 'actions',
              title: <Text mr="xs">Row actions</Text>,
              textAlignment: 'right',
              render: (data: User) => (
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
          datas={users || []}
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


