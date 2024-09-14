import "@@/common/styles/globals.css";
import DashboardRoot from "@@/views/Dashboard";
import ApiKeysHome from "@@/views/Dashboard/ApiKeysPage";
import ApplicationPage from "@@/views/Dashboard/ApplicationPage";
import ApplicationsPage from "@@/views/Dashboard/ApplicationsPage";
import DashboardHomePage from "@@/views/Dashboard/DashboardHomePage";
import EventTypePage from "@@/views/Dashboard/EventTypePage";
import EventTypesPage from "@@/views/Dashboard/EventTypesPage";
import SignInPage from "@@/views/signin";
import SignUpPage from "@@/views/signup";
import { Theme } from "@radix-ui/themes";
import "@radix-ui/themes/styles.css";
import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import { HelmetProvider } from "react-helmet-async";
import {
  createBrowserRouter,
  Navigate,
  RouterProvider,
} from "react-router-dom";
import { config } from "./config";

const basePath = config.basePath;

const router = createBrowserRouter([
  { path: `${basePath}`, element: <Navigate to={`${basePath}/signin`} /> },
  { path: `${basePath}/signin`, element: <SignInPage /> },
  { path: `${basePath}/signup`, element: <SignUpPage /> },
  {
    path: `${basePath}/:environment`,
    element: <DashboardRoot />,
    children: [
      {
        path: "",
        element: <DashboardHomePage />,
      },
      {
        path: "applications",
        element: <ApplicationsPage />,
      },
      {
        path: "applications/:appId",
        element: <ApplicationPage />,
      },
      {
        path: "event-types",
        element: <EventTypesPage />,
      },
      {
        path: "event-types/:eventTypeId",
        element: <EventTypePage />,
      },
      {
        path: "api-keys",
        element: <ApiKeysHome />,
      },
    ],
  },
]);

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <HelmetProvider>
      <Theme accentColor="indigo" appearance="dark">
        <RouterProvider router={router} />
      </Theme>
    </HelmetProvider>
  </StrictMode>
);
