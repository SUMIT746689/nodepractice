import { Layout, Menu } from "antd";
import { Form, Outlet, redirect, useNavigate } from "react-router-dom";
import { Button } from "../../components/ui/button";
import { ofetch } from "@/lib/ofetch";
import { useState } from "react";
import { Head, NavbarNested } from "./nav/Navbar";
import { AppShell, Navbar, Header } from '@mantine/core';

export default function MainLayout() {
  return (
    <>
      {/* <AppShell
        navbar={<NavbarNested />}
        header={<Header><Head/></Header>}
      >
        <Outlet />
      </AppShell> */}
      <div className="flex">
        <NavbarNested />
        <Outlet/>
      </div>
    </>
  )
}


export const LayoutRouteAction = async ({ request }) => {
  console.log("yes");

  try {
    await ofetch("/logout", { method: "post" })
    document.cookie = "Authorization="
    return redirect("/login")

  } catch (error) {

  }

  return null
}
