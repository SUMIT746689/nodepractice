
import { PageHeaderWrapper } from "@/components/PageHeaderWrapper";
import CreateOrUpdateData from "@/content/companies/CreateOrUpdateData";
import ShowData from "@/content/companies/ShowData";
import { User } from "@/types/users";
import { useState } from "react";


export default function Companies() {

  const [editUser, setEditUser] = useState<User>();

  const addEditUser = (user: User): void => {
    setEditUser(() => user)
  }

  return (
    <>
      <PageHeaderWrapper name="Companies">
        <CreateOrUpdateData editData={editUser} setEditData={setEditUser} />
      </PageHeaderWrapper>

      <div className="px-6 py-3 w-full min-h-fit">
        <ShowData addEditData={addEditUser} />
      </div>
    </>
  )
}


