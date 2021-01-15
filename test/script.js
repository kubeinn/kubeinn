import http from 'k6/http';

import { sleep } from 'k6';

export default function () {

    http.get('http://51.222.35.240:30050/');

    sleep(1);

}