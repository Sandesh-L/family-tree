
interface Member{
    id: number;
    firstName: string;
    middleName?: string;
    lastName: string;
    parents: number[];
    spouses: number[];
    children: number[];
}

export class MemberNode {
    Id: number;
    firstName: string;
    middleName?: string;
    lastName: string;
    parents: number[];
    spouses: number[];
    children: number[];
    
    constructor(member: Member){
        this.Id = member.id
        this.firstName = member.firstName;
        this.middleName = member.middleName;
        this.lastName = member.lastName;
        this.parents = member.parents;
        this.spouses = member.spouses;
        this.children = member.children;
    }
}