"use client";

import React, { useEffect, useState } from "react";
import { getUsers } from "../../../services/api";

const UsersPage = () => {
    const [users, setUsers] = useState([]);

    useEffect(() => {
        const fetchUsers = async () => {
            try {
                const data = await getUsers();
                setUsers(data);
            } catch (error) {
                console.error('Error loading users:', error);
            }
        };

        fetchUsers();
    }, []);

    return (
        <div>
            <h1>User Profiles</h1>
            <ul>
                {users.map((user) => (
                    <li key={user.id}>
                        <h2>{user.name}</h2>
                        <p>{user.email}</p>
                        <p>{user.bio}</p>
                        <img src={user.profile_picture} alt={user.name} />
                    </li>
                ))}
            </ul>
        </div>
    );
};

export default UsersPage;
