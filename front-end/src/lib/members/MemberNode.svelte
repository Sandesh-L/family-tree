<script>
// @ts-nocheck
    export let memberNode = null;
</script>

<div class="member-node">
    <div class="member-name">
        {memberNode.firstName} {memberNode.lastName}
    </div>
    <div class="member-details">
        <p><strong>MiddleName:</strong> {memberNode.middleName} </p>
    </div>
    
</div>

<style>
    .member-node {
        border: 2px solid #333;
        padding: 16px;
        margin: 8px;
        border-radius: 8px;
        width: 200px;
        text-align: center;
        box-shadow: 2px 2px 12px rgba(0, 0, 0, 0.1);
    }
    .member-name {
        font-size: 1.2rem;
        font-weight: bold;
        margin-bottom: 8px;
    }
    .member-details {
        font-size: 0.9rem;
        color: #555;
    }
</style>