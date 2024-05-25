import {
  BottomNavigation,
  BottomNavigationAction,
  IconButton,
  Tooltip,
} from "@mui/material";
import GitHubIcon from "@mui/icons-material/GitHub";
import LinkedInIcon from "@mui/icons-material/LinkedIn";

function Footer({ position = "fixed" }) {
  return (
    <BottomNavigation
      showLabels={false}
      sx={{ position: { position }, bottom: 0, left: 0, right: 0 }}
    >
      <Tooltip title="GitHub Repository">
        <BottomNavigationAction
          label="GitHub Repository"
          icon={
            <IconButton
              href="https://github.com/ckng0221/library-app"
              target="_blank"
            >
              <GitHubIcon />
            </IconButton>
          }
        />
      </Tooltip>

      <Tooltip title="Author's Linkedlin">
        <BottomNavigationAction
          label="Author's Linkedlin"
          icon={
            <IconButton
              href="https://www.linkedin.com/in/choon-khon-ng-7b5755149/"
              target="_blank"
            >
              <LinkedInIcon />
            </IconButton>
          }
        />
      </Tooltip>
    </BottomNavigation>
  );
}

export default Footer;
