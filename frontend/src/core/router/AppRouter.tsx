// src/core/router/AppRouter.tsx
import {BrowserRouter, Route, Routes} from "react-router-dom";
import TrainsPage from "../../modules/trains/pages/TrainsPage";
import TrainDetailPage from "../../modules/trains/pages/TrainDetailPage";
import CreateTrainPage from "../../modules/trains/pages/CreateTrainPage.tsx";
import IndexPage from "../../modules/trains/pages/IndexPage.tsx";
import UsersPage from "../../modules/trains/pages/UsersPage.tsx";
import CreateUserPage from "../../modules/trains/pages/CreateUserPage.tsx";
import UserDetailPage from "../../modules/trains/pages/UserDetailPage.tsx";
import StationsPage from "../../modules/trains/pages/StationsPage.tsx";
import CreateStationPage from "../../modules/trains/pages/CreateStationPage.tsx";
import EditStationPage from "../../modules/trains/pages/EditStationPage.tsx";
import RoutesPage from "../../modules/trains/pages/RoutesPage.tsx";
import CreateRoutePage from "../../modules/trains/pages/CreateRoutePage.tsx";
import EditRoutePage from "../../modules/trains/pages/EditRoutePage.tsx";

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

                <Route path="/routes" element={<RoutesPage />} />
                <Route path="/routes/create" element={<CreateRoutePage />} />
                <Route path="/routes/:id" element={<EditRoutePage />} />
            </Routes>
        </BrowserRouter>
    );
}
