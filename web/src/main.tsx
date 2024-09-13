import "@@/common/styles/globals.css";
import ApiKeysHome from "@@/pages/Dashboard/ApiKeysPage.tsx";
import ApplicationPage from "@@/pages/Dashboard/ApplicationPage.tsx";
import ApplicationsPage from "@@/pages/Dashboard/ApplicationsPage.tsx";
import DashboardHomePage from "@@/pages/Dashboard/DashboardHomePage.tsx";
import EventTypePage from "@@/pages/Dashboard/EventTypePage.tsx";
import EventTypesPage from "@@/pages/Dashboard/EventTypesPage.tsx";
import DashboardRoot from "@@/pages/Dashboard/index.tsx";
import SignInPage from "@@/pages/signin.tsx";
import SignUpPage from "@@/pages/signup.tsx";
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
