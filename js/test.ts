interface Message {
    name: string;
    headers: Header[];
}

interface Header {
    name: string;
    value: string;
}

interface BodyPart {
    name: string;
    data: Int8Array;
}
