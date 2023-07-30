
import CreateOrUpdateData from "@/content/users/CreateOrUpdateData";
import ShowData from "@/content/users/ShowData";
import { User } from "@/types/users";
import { Header, Text } from "@mantine/core";
import { useState } from "react";


export default function UserIndex() {

  const [editUser, setEditUser] = useState<User>();

  const addEditUser = (user: User):void =>{
    setEditUser(()=>user)
  }

  return (
    <>
      <Header height={64} display="flex" sx={{ justifyContent: 'space-between' }}>
        <Text pt={18} pl={18} size={"lg"} weight={600} color="orange">Users</Text>
        <CreateOrUpdateData editData={editUser} setEditData={setEditUser} />
      </Header>

      <div className="px-6 py-3 w-full min-h-fit">
        <ShowData addEditData={addEditUser}/>
      </div>
    </>
  )
}


