import DeleteIcon from "@mui/icons-material/Delete";
import EditIcon from "@mui/icons-material/Edit";
import {
  Button,
  Grid,
  IconButton,
  List,
  ListItem,
  ListItemText,
  TextField,
  Tooltip,
  Typography,
} from "@mui/material";
import { Dispatch, SetStateAction, useEffect, useState } from "react";
import { getUserAddresses, updateUserById } from "../api/user";
import BasicSelect from "../components/BasicSelect";
import { IAddress, IUser } from "../interfaces/user";
import toast from "react-hot-toast";
import AddCircleIcon from "@mui/icons-material/AddCircle";

export default function Profile({
  user,
  setUser,
}: {
  user: IUser;
  setUser: Dispatch<SetStateAction<IUser>>;
}) {
  const [addresses, setAddresses] = useState<IAddress[]>([
    {
      id: 0,
      street: "",
      city: "",
      state: "",
      zip: "",
    },
  ]);
  const [userProfile, setUserProfile] = useState({
    name: user.name,
    default_address_id: user.default_address_id,
  });

  useEffect(() => {
    async function loadData() {
      if (!user.id) return;
      // console.log(user);

      setUserProfile({
        name: user.name,
        default_address_id: user.default_address_id,
      });
      const res = await getUserAddresses(user.id);
      if (res?.data.status === "success") {
        setAddresses(res.data?.data);
      }
    }
    loadData();
  }, [user]);

  function handleAddressChange(event: any) {
    setUserProfile({
      ...event.target.value,
      default_address_id: event.target.value,
    });
  }

  function handleDeleteAddress(addressId: string | number | undefined) {
    if (!addressId) return;
    //TODO: delete address
    confirm("are you sure?");
  }

  async function handleUpdate() {
    setUser({
      ...user,
      name: userProfile.name,
      default_address_id: userProfile.default_address_id,
    });
    const res = await updateUserById(user.id, {
      name: userProfile.name,
      default_address_id: userProfile.default_address_id,
    });
    if (res?.status === 200) {
      toast.success("Updated profile!");
    } else {
      toast.error("Failed to update profile");
    }
  }

  return (
    <div className="justify-start">
      <TextField
        label="Name"
        value={userProfile.name}
        className="mb-4"
        onChange={(e) => {
          setUserProfile({ ...userProfile, name: e.target.value });
        }}
      />
      <div className="mb-8 justify-start">
        <Typography align="left" className="mb-4">
          Default Address
        </Typography>
        <div className="mb-4">
          <BasicSelect
            options={addresses}
            value={userProfile.default_address_id}
            label="Address"
            handleChange={handleAddressChange}
          />
        </div>
        <div className="flex justify-end">
          <Button variant="contained" onClick={handleUpdate}>
            Update
          </Button>
        </div>
      </div>
      <Grid item xs={12} md={6}>
        <div className="flex gap-2 content-center">
          <Typography
            sx={{ mt: 4, mb: 2 }}
            align="left"
            variant="h6"
            component="div"
          >
            Addresses:
          </Typography>
          <Tooltip title="add new address" className="">
            <IconButton edge="end" aria-label="add">
              <AddCircleIcon color="primary" />
            </IconButton>
          </Tooltip>
        </div>

        {addresses.map((address) => (
          <List key={address.id} dense={true}>
            <ListItem
              secondaryAction={
                <>
                  <IconButton edge="end" aria-label="edit">
                    <EditIcon />
                  </IconButton>
                  <IconButton
                    edge="end"
                    aria-label="delete"
                    onClick={() => {
                      handleDeleteAddress(address?.id);
                    }}
                  >
                    <DeleteIcon />
                  </IconButton>
                </>
              }
            >
              <ListItemText
                primary={
                  <>
                    <div>{address.street}</div>
                    <div>{address.city}</div>
                    <div>{address.street}</div>
                    <div>{address.city}</div>
                    <div>{address.zip}</div>
                  </>
                }
              />
            </ListItem>
          </List>
        ))}
      </Grid>
    </div>
  );
}
