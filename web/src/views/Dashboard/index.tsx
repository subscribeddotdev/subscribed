import { LayoutDashboard } from "@@/common/layouts/LayoutDashboard/LayoutDashboard";
import { paths } from "@@/constants";
import { retrieveTokenFromTheClient } from "@@/modules/Auth/token";
import { Spinner } from "@radix-ui/themes";
import { useEffect, useState } from "react";
import { Outlet, useNavigate } from "react-router-dom";

export default function DashboardRoot() {
  const [ready, setReady] = useState(false);

  const navigate = useNavigate();
  useEffect(() => {
    if (!retrieveTokenFromTheClient()) {
      navigate(paths.signin);
      return;
    }

    setReady(true);
  }, [navigate]);

  if (!ready) {
    return <Spinner />;
  }

  return (
    <LayoutDashboard>
      <Outlet />
    </LayoutDashboard>
  );
}
