-- Resources available
CREATE TABLE IF NOT EXISTS rbac_resource(
  id CHAR(31) PRIMARY KEY,
  name VARCHAR(32),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP
);

-- Actions available on resources
CREATE TABLE IF NOT EXISTS rbac_action(
  id CHAR(31) PRIMARY KEY,
  name VARCHAR(32),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP
);

-- Roles that a user may have
CREATE TABLE IF NOT EXISTS rbac_role(
  id CHAR(31) PRIMARY KEY,
  name VARCHAR(32),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP
);

-- Join table to enable many-to-many relationship between roles and actions on resources
CREATE TABLE IF NOT EXISTS rbac_permission(
  role_id CHAR(31) NOT NULL,
  action_id CHAR(31) NOT NULL,
  resource_id CHAR(31) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP,
  PRIMARY KEY (role_id, action_id, resource_id),
  FOREIGN KEY (role_id) REFERENCES rbac_role(id),
  FOREIGN KEY (action_id) REFERENCES rbac_action(id),
  FOREIGN KEY (resource_id) REFERENCES rbac_resource(id)
);

-- Join table to enable many-to-many relationship between users and roles
CREATE TABLE IF NOT EXISTS rbac_app_user_role(
  id SERIAL PRIMARY KEY,
  app_user_id CHAR(31) NOT NULL,
  role_id CHAR(31) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP,
  FOREIGN KEY (app_user_id) REFERENCES app_user(id) DEFERRABLE INITIALLY DEFERRED,
  FOREIGN KEY (role_id) REFERENCES rbac_role(id) DEFERRABLE INITIALLY DEFERRED
);