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
import TicketsPage from "../../modules/tickets/pages/TicketsPage.tsx";
import EditTicketPage from "../../modules/tickets/pages/EditTicketPage.tsx";
import NotificationsPage from "../../modules/subscriptions/pages/SubscriptionsPage.tsx";
import CreateSubscriptionPage from "../../modules/subscriptions/pages/CreateSubscriptionPage.tsx";

export default function AppRouter() {
    return (
        <BrowserRouter>
            <Routes>
                <Route path="/" element={<IndexPage />} />
                <Route path="/trains" element={<TrainsPage />} />
                <Route path="/admin/trains/:uuid" element={<TrainDetailPage />} />
                <Route path="/trains/create" element={<CreateTrainPage />} />
                <Route path="/users" element={<UsersPage />} />
                <Route path="/users/create" element={<CreateUserPage />} />
                <Route path="/users/:id" element={<UserDetailPage />} />
                <Route path="/stations" element={<StationsPage />} />
                <Route path="/stations/create" element={<CreateStationPage />} />
                <Route path="/stations/:id" element={<EditStationPage />} />

                <Route path="/schedules" element={<SchedulesPage />} />
                <Route path="/schedules/create" element={<CreateSchedulePage />} />
                <Route path="/schedules/:id" element={<EditSchedulePage />} />
                <Route path="/schedules/:id/stops" element={<ScheduleStopsPage />} />

                <Route path="/tickets" element={<TicketsPage />} />
                <Route path="/tickets/:uuid" element={<EditTicketPage />} />

                <Route path="/subscriptions" element={<NotificationsPage />} />
                <Route path="/subscriptions/create" element={<CreateSubscriptionPage />} />

            </Routes>
        </BrowserRouter>
    );
}
