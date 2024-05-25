import AddCircleIcon from "@mui/icons-material/AddCircle";
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
import Dialog from "@mui/material/Dialog";
import DialogActions from "@mui/material/DialogActions";
import DialogContent from "@mui/material/DialogContent";
import DialogTitle from "@mui/material/DialogTitle";
import { useTheme } from "@mui/material/styles";
import useMediaQuery from "@mui/material/useMediaQuery";
import {
  Dispatch,
  Fragment,
  SetStateAction,
  useEffect,
  useReducer,
  useState,
} from "react";
import toast from "react-hot-toast";
import {
  createAddress,
  deleteAddressById,
  getUserAddresses,
  updateAddressById,
  updateUserById,
} from "../api/user";
import BasicSelect from "../components/BasicSelect";
import { IAddress, IUser } from "../interfaces/user";

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
  const [openForm, setOpenForm] = useState(false);
  const emptyAddress: IAddress = {
    street: "",
    state: "",
    city: "",
    zip: "",
    user_id: user.id,
  };
  const [formAddress, setFormAddress] = useState<IAddress>(emptyAddress);
  const [newAddress, setNewAddress] = useState(false);
  const [state, setState] = useState(0);
  const [errorFields, setErrorFields] = useState<string[]>([]);

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
  }, [user, state]);

  function handleAddressChange(event: any) {
    console.log(event.target.value);

    setUserProfile((prev) => {
      return {
        ...prev,
        default_address_id: event.target.value,
      };
    });
  }

  async function handleDeleteAddress(addressId: string | number | undefined) {
    if (!addressId) return;
    //TODO: replace with dialog
    const isConfirm = confirm("are you sure?");
    if (!isConfirm) return;

    const res = await deleteAddressById(String(addressId));
    if (res?.status) {
      toast.success("Removed address!");
      setState((prev) => prev + 1);
    } else {
      toast.error("Failed to remove address");
    }
  }

  async function handleUpdate() {
    console.log(userProfile);
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

  async function handleEditAdress(address: IAddress) {
    setErrorFields([]);
    setNewAddress(false);
    setFormAddress(address);
    setOpenForm((prev) => !prev);
  }

  async function handleAddNewAddress() {
    setErrorFields([]);
    setNewAddress(true);
    setFormAddress(emptyAddress);
    setOpenForm((prev) => !prev);
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
            <IconButton
              edge="end"
              aria-label="add"
              onClick={handleAddNewAddress}
            >
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
                    <EditIcon onClick={() => handleEditAdress(address)} />
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
                    <div>{address.state}</div>
                    <div>{address.zip}</div>
                  </>
                }
              />
            </ListItem>
          </List>
        ))}
      </Grid>
      <AddressForm
        open={openForm}
        setOpen={setOpenForm}
        formAddress={formAddress}
        setFormAddress={setFormAddress}
        setState={setState}
        isNew={newAddress}
        errorFields={errorFields}
        setErrorFields={setErrorFields}
      />
    </div>
  );
}

function AddressForm({
  open,
  setOpen,
  isNew,
  formAddress,
  setFormAddress,
  setState,
  errorFields,
  setErrorFields,
}: {
  open: boolean;
  setOpen: Dispatch<SetStateAction<boolean>>;
  isNew: boolean;
  formAddress: IAddress;
  setFormAddress: Dispatch<SetStateAction<IAddress>>;
  setState: Dispatch<SetStateAction<number>>;
  errorFields: string[];
  setErrorFields: Dispatch<SetStateAction<string[]>>;
}) {
  const theme = useTheme();
  const fullScreen = useMediaQuery(theme.breakpoints.down("md"));

  const handleClose = () => {
    setOpen(false);
  };

  async function handleUpdate() {
    if (!formAddress.id) return;
    if (!validateForm(formAddress)) return;

    const res = await updateAddressById(String(formAddress.id), formAddress);
    if (res?.status) {
      toast.success("Updated address!");
      setOpen(false);
      setState((prev) => prev + 1);
    } else {
      toast.error("Failed to updated address");
    }
  }
  async function handleSave() {
    if (!validateForm(formAddress)) return;

    const res = await createAddress(formAddress);
    if (res?.status) {
      toast.success("Added address!");
      setOpen(false);
      setState((prev) => prev + 1);
    } else {
      toast.error("Failed to add address");
    }
  }

  function validateForm(formAddress: IAddress) {
    setErrorFields([]);

    let errorFieldCount = 0;
    for (const [key, value] of Object.entries(formAddress)) {
      if (value == "") {
        setErrorFields((prev) => [...prev, key]);
        errorFieldCount++;
      }
    }
    if (errorFieldCount > 0) {
      return false;
    }
    return true;
  }

  return (
    <Fragment>
      <Dialog
        fullScreen={fullScreen}
        open={open}
        onClose={handleClose}
        aria-labelledby="responsive-dialog-title"
      >
        <DialogTitle id="responsive-dialog-title">
          {isNew ? "New " : "Edit "}Address
        </DialogTitle>
        <DialogContent>
          <div className="grid grid-cols-3 gap-2">
            <TextField
              label="Street"
              value={formAddress.street}
              className="mb-4 col-span-3"
              onChange={(e) => {
                setFormAddress({ ...formAddress, street: e.target.value });
              }}
              multiline
              variant="standard"
              required
              error={errorFields.includes("street")}
              helperText={
                errorFields.includes("street") ? "Street cannot be empty" : ""
              }
            />
            <TextField
              label="Zip"
              value={formAddress.zip}
              className="mb-4"
              onChange={(e) => {
                setFormAddress({ ...formAddress, zip: e.target.value });
              }}
              variant="standard"
              required
              error={errorFields.includes("zip")}
              inputProps={{ maxLength: 5 }}
              helperText={
                errorFields.includes("zip") ? "Zip cannot be empty" : ""
              }
            />
            <TextField
              label="City"
              value={formAddress.city}
              className="mb-4"
              onChange={(e) => {
                setFormAddress({ ...formAddress, city: e.target.value });
              }}
              variant="standard"
              error={errorFields.includes("city")}
              required
              helperText={
                errorFields.includes("city") ? "City cannot be empty" : ""
              }
            />
            <TextField
              label="State"
              value={formAddress.state}
              className="mb-4"
              onChange={(e) => {
                setFormAddress({ ...formAddress, state: e.target.value });
              }}
              variant="standard"
              error={errorFields.includes("state")}
              required
              helperText={
                errorFields.includes("state") ? "State cannot be empty" : ""
              }
            />
          </div>
        </DialogContent>
        <DialogActions>
          <Button autoFocus onClick={handleClose}>
            Cancel
          </Button>
          {isNew ? (
            <Button onClick={handleSave} autoFocus>
              Save
            </Button>
          ) : (
            <Button onClick={handleUpdate} autoFocus>
              Update
            </Button>
          )}
        </DialogActions>
      </Dialog>
    </Fragment>
  );
}
