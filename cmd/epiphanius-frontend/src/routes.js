import Vue from "vue";
import Router from "vue-router";

Vue.use(Router);

const router = new Router({
    mode: 'history',
    routes: [
        {
            path: "/",
            name: "Holidays",
            component: () => import("./components/Holidays"),
        },
        {
            path: "/holidays",
            name: "Holidays",
            component: () => import("./components/Holidays"),
        },
        {
            path: "/holiday/:id",
            name: "Holiday",
            component: () => import("./components/Holiday"),
        },
    ]
});

export default router;