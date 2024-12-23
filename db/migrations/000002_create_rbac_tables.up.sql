-- Roles that a user may have
CREATE TABLE IF NOT EXISTS rbac_role(
  id CHAR(24) PRIMARY KEY,
  name VARCHAR(32),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP
);

-- Permissions that roles can have
CREATE TABLE IF NOT EXISTS rbac_permission(
  id CHAR(24) PRIMARY KEY,
  name VARCHAR(32),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP
);

-- Join table to enable many-to-many relationship between roles and permissions
CREATE TABLE IF NOT EXISTS rbac_roles_permissions(
  id SERIAL PRIMARY KEY,
  role_id CHAR(24) NOT NULL,
  permission_id CHAR(24) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP,
  FOREIGN KEY (role_id) REFERENCES rbac_role(id),
  FOREIGN KEY (permission_id) REFERENCES rbac_permission(id)
);

-- Join table to enable many-to-many relationship between users and roles
CREATE TABLE IF NOT EXISTS rbac_app_users_roles(
  id SERIAL PRIMARY KEY,
  app_user_id CHAR(24) NOT NULL,
  role_id CHAR(24) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP,
  FOREIGN KEY (app_user_id) REFERENCES app_user(id),
  FOREIGN KEY (role_id) REFERENCES rbac_role(id)
);