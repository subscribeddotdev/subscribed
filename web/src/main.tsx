import "@@/common/styles/globals.css";
import DashboardRoot from "@@/Pages/Dashboard";
import ApiKeysHome from "@@/Pages/Dashboard/ApiKeysPage";
import ApplicationPage from "@@/Pages/Dashboard/ApplicationPage";
import ApplicationsPage from "@@/Pages/Dashboard/ApplicationsPage";
import DashboardHomePage from "@@/Pages/Dashboard/DashboardHomePage";
import EventTypePage from "@@/Pages/Dashboard/EventTypePage";
import EventTypesPage from "@@/Pages/Dashboard/EventTypesPage";
import SignInPage from "@@/Pages/signin";
import SignUpPage from "@@/Pages/signup";
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

const router = createBrowserRouter([
  { path: "/", element: <Navigate to="/signin" /> },
  { path: "/signin", element: <SignInPage /> },
  { path: "/signup", element: <SignUpPage /> },
  {
    path: "/:environment",
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
