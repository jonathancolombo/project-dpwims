import {BrowserRouter, Route, Routes} from "react-router-dom";
import TrainsPage from "../../modules/trains/pages/TrainsPage";
import TrainDetailPage from "../../modules/trains/pages/TrainDetailPage";
import CreateTrainPage from "../../modules/trains/pages/CreateTrainPage.tsx";
import IndexPage from "../pages/IndexPage.tsx";
import UsersPage from "../../modules/users/pages/UsersPage.tsx";
import CreateUserPage from "../../modules/users/pages/CreateUserPage.tsx";
import UserDetailPage from "../../modules/users/pages/UserDetailPage.tsx";
import StationsPage from "../../modules/trains/pages/StationsPage.tsx";
import CreateStationPage from "../../modules/trains/pages/CreateStationPage.tsx";
import EditStationPage from "../../modules/trains/pages/EditStationPage.tsx";
import SchedulesPage from "../../modules/trains/pages/SchedulesPage.tsx";
import CreateSchedulePage from "../../modules/trains/pages/CreateSchedulePage.tsx";
import EditSchedulePage from "../../modules/trains/pages/EditSchedulePage.tsx";
import ScheduleStopsPage from "../../modules/trains/pages/ScheduleStopsPage.tsx";
import EditTicketPage from "../../modules/tickets/pages/EditTicketPage.tsx";
import NotificationsPage from "../../modules/subscriptions/pages/SubscriptionsPage.tsx";
import CreateSubscriptionPage from "../../modules/subscriptions/pages/CreateSubscriptionPage.tsx";
import CreateTicketPage from "../../modules/tickets/pages/CreateTicketPage.tsx";
import CreatePaymentPage from "../../modules/tickets/pages/CreatePaymentPage.tsx";
import TransactionsPage from "../../modules/tickets/pages/TransactionsPage.tsx";
import {RequireAdmin} from "../../modules/authentication/pages/RequireAdmin.tsx";
import LoginPage from "../../modules/authentication/pages/LoginPage.tsx";
import {RequireUser} from "../../modules/authentication/pages/RequireUser.tsx";
import UserHomePage from "../../modules/authentication/pages/UserHomePage.tsx";
import MyTicketsPage from "../../modules/tickets/pages/MyTicketsPage";
import AdminHomePage from "../../modules/authentication/pages/AdminHomePage";
import UserSchedulesAndTicketsPage from "../../modules/users/pages/UserSchedulesAndTicketsPage.tsx";
import UserSubscriptionsPage from "../../modules/subscriptions/pages/UserSubscriptionsPage.tsx";

export default function AppRouter() {
    return (
        <BrowserRouter>
            <Routes>

                {/* HOME + LOGIN */}
                <Route path="/" element={<IndexPage />} />
                <Route path="/login" element={<LoginPage />} />

                {/* AREA UTENTE */}
                <Route
                    path="/user"
                    element={
                        <RequireUser>
                            <UserHomePage />
                        </RequireUser>
                    }
                />

                <Route
                    path="/user/schedules"
                    element={
                        <RequireUser>
                            <UserSchedulesAndTicketsPage />
                        </RequireUser>
                    }
                />

                <Route
                    path="/my-tickets"
                    element={
                        <RequireUser>
                            <MyTicketsPage />
                        </RequireUser>
                    }
                />

                <Route
                    path="/user/tickets"
                    element={
                        <RequireUser>
                            <MyTicketsPage />
                        </RequireUser>
                    }
                />

                <Route
                    path="/user/subscriptions"
                    element={
                        <RequireUser>
                            <UserSubscriptionsPage />
                        </RequireUser>
                    }
                />

                <Route
                    path="/user/buy-ticket/:scheduleId"
                    element={
                        <RequireUser>
                            <UserSchedulesAndTicketsPage />
                        </RequireUser>
                    }
                />

                <Route
                    path="/buy-ticket/:scheduleId"
                    element={
                        <RequireUser>
                            <UserSchedulesAndTicketsPage />
                        </RequireUser>
                    }
                />


                {/* AREA ADMIN */}
                <Route
                    path="/admin"
                    element={
                        <RequireAdmin>
                            <AdminHomePage />
                        </RequireAdmin>
                    }
                />

                <Route
                    path="/trains"
                    element={
                        <RequireAdmin>
                            <TrainsPage />
                        </RequireAdmin>
                    }
                />

                <Route
                    path="/admin/trains/:uuid"
                    element={
                        <RequireAdmin>
                            <TrainDetailPage />
                        </RequireAdmin>
                    }
                />

                <Route
                    path="/trains/create"
                    element={
                        <RequireAdmin>
                            <CreateTrainPage />
                        </RequireAdmin>
                    }
                />

                <Route
                    path="/users"
                    element={
                        <RequireAdmin>
                            <UsersPage />
                        </RequireAdmin>
                    }
                />

                <Route
                    path="/users/create"
                    element={
                        <RequireAdmin>
                            <CreateUserPage />
                        </RequireAdmin>
                    }
                />

                <Route
                    path="/users/:id"
                    element={
                        <RequireAdmin>
                            <UserDetailPage />
                        </RequireAdmin>
                    }
                />

                <Route
                    path="/stations"
                    element={
                        <RequireAdmin>
                            <StationsPage />
                        </RequireAdmin>
                    }
                />

                <Route
                    path="/stations/create"
                    element={
                        <RequireAdmin>
                            <CreateStationPage />
                        </RequireAdmin>
                    }
                />

                <Route
                    path="/stations/:id"
                    element={
                        <RequireAdmin>
                            <EditStationPage />
                        </RequireAdmin>
                    }
                />

                <Route
                    path="/schedules"
                    element={
                        <RequireAdmin>
                            <SchedulesPage />
                        </RequireAdmin>
                    }
                />

                <Route
                    path="/schedules/create"
                    element={
                        <RequireAdmin>
                            <CreateSchedulePage />
                        </RequireAdmin>
                    }
                />

                <Route
                    path="/schedules/:id"
                    element={
                        <RequireAdmin>
                            <EditSchedulePage />
                        </RequireAdmin>
                    }
                />

                <Route
                    path="/schedules/:id/stops"
                    element={
                        <RequireAdmin>
                            <ScheduleStopsPage />
                        </RequireAdmin>
                    }
                />

                <Route
                    path="/transactions"
                    element={
                        <RequireAdmin>
                            <TransactionsPage />
                        </RequireAdmin>
                    }
                />

                <Route
                    path="/tickets/create"
                    element={
                        <RequireAdmin>
                            <CreateTicketPage />
                        </RequireAdmin>
                    }
                />

                <Route
                    path="/tickets/:uuid"
                    element={
                        <RequireAdmin>
                            <EditTicketPage />
                        </RequireAdmin>
                    }
                />

                <Route
                    path="/subscriptions"
                    element={
                        <RequireAdmin>
                            <NotificationsPage />
                        </RequireAdmin>
                    }
                />

                <Route
                    path="/subscriptions/create"
                    element={
                        <RequireAdmin>
                            <CreateSubscriptionPage />
                        </RequireAdmin>
                    }
                />

                <Route
                    path="/payments/create"
                    element={
                        <RequireAdmin>
                            <CreatePaymentPage />
                        </RequireAdmin>
                    }
                />

            </Routes>
        </BrowserRouter>
    );
}