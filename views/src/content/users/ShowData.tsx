import PaginationWithPageSizeSelectorWrapper from "@/components/PaginationWithPageSizeSelectorWrapper";
import { useDeleteUserMutation, useGetAllUsersQuery } from "@/redux/services/user";
import { User } from "@/types/users";
import { ActionIcon, Group, Text } from "@mantine/core";
import { notifications } from "@mantine/notifications";
import { IconEdit, IconTrash } from "@tabler/icons-react";
import React from "react";


interface ShowDataInterface {
  // setEditUser: React.Component<User,User>
  setEditUser: any;
} 

const ShowData:React.FC<ShowDataInterface> =({ setEditUser })=> {

  const { data: users } = useGetAllUsersQuery();
  const [deleteUser] = useDeleteUserMutation();
  console.log({users})
  const handleDelete = (id: number): void => {
    deleteUser(id).unwrap()
      .then(() => { notifications.show({ message: 'deleted successfully', }) })
      .catch((error: { message: string }) => { notifications.show({ message: error.message, color: 'red' }) })
  }

  return (
    <>
      <PaginationWithPageSizeSelectorWrapper
        headColumns={[
          { accessor: 'id', width: 100 },
          { accessor: 'first_name' },
          { accessor: 'last_name' },
          { accessor: 'username' },
          { accessor: 'email' },
          // { accessor: 'role.title' },
          { accessor: 'phone_number' },
          {
            accessor: 'actions',
            title: <Text mr="xs">Row actions</Text>,
            textAlignment: 'right',
            render: (data: User) => (
              <Group spacing={4} position="right" noWrap>
                {/* // eslint-disable-next-line @typescript-eslint/no-unsafe-return */}
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
    </>
  )
}

export default ShowData;