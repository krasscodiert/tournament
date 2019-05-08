import "spectre.css/dist/spectre.min.css";
import "spectre.css/dist/spectre-exp.min.css";
import "spectre.css/dist/spectre-icons.min.css";

import "./style.scss";

import m from "mithril";

import { HomeView } from "./views/HomeView";
import { TournamentsView } from "./views/TournamentsView";
import { CreateTournamentView } from "./views/CreateTournamentView";

m.route(document.body, "/home", {
    "/home" : HomeView,
    "/tournaments" : TournamentsView,
    "/create-tournament": CreateTournamentView
});