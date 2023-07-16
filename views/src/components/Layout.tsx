import { Layout, Menu } from "antd";
import { Form, Outlet, redirect, useNavigate } from "react-router-dom";
import { Button } from "./ui/button";
import { ofetch } from "@/lib/ofetch";
import { useState } from "react";

export default function MainLayout() {
  return (
    <>
      <div>Layout</div>
      <div>{<Outlet/>}</div>
    </>
  )
}


export const LayoutRouteAction = async ({ request }) => {
  console.log("yes");
  
  try {
    await ofetch("/logout", {method: "post"})
    document.cookie = "Authorization="
    return redirect("/login")
    
  } catch (error) {
    
  }

  return null
}
